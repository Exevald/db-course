package main

import (
	"context"
	"net/http"
	"orgchart/pkg/integrationaltests/tests"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"orgchart/pkg/orgchart/common/server"
)

const (
	waitTimeout = 30 * time.Second
)

type Config struct {
	DBHost               string `envconfig:"db_host"`
	DBName               string `envconfig:"db_name"`
	DBUser               string `envconfig:"db_user"`
	DBPassword           string `envconfig:"db_password"`
	DBMaxConn            int    `envconfig:"db_max_conn" default:"0"`
	DBConnectionLifetime int    `envconfig:"db_conn_lifetime" default:"0"`

	ServiceHTTPEndpoint string `envconfig:"service_http_endpoint" default:"http://orgchart-dev-for-tests:8082"`
	ServeRESTAddress    string `envconfig:"serve_rest_address" default:":8082"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var httpServer *http.Server
	defer cancel()

	var config Config
	if err := envconfig.Process("orgchart", &config); err != nil {
		log.WithError(err).Fatal("Failed to process env vars")
	}

	log.Info("Waiting for services to start")
	if err := waitForServices(ctx, config, waitTimeout); err != nil {
		log.WithError(err).Fatal("Failed to wait for services to start")
	}

	errChan := make(chan struct{})
	serverHub := server.NewHub(errChan)
	serverHub.Serve(func() error {
		router := mux.NewRouter()
		httpServer = &http.Server{
			Handler:           router,
			Addr:              config.ServeRESTAddress,
			ReadHeaderTimeout: 10 * time.Second,
			ReadTimeout:       time.Hour,
			WriteTimeout:      time.Hour,
		}
		return httpServer.ListenAndServe()
	}, func() error {
		cancel()
		return httpServer.Shutdown(context.Background())
	})
	tests.RunTests()
}

func waitForServices(ctx context.Context, config Config, timeout time.Duration) error {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return waitForServiceReady(ctx, config.ServiceHTTPEndpoint)
	})
	return eg.Wait()
}

func waitForServiceReady(ctx context.Context, endpoint string) error {
	const readyPath = "/resilience/ready"
	request, err := http.NewRequest(http.MethodGet, endpoint+readyPath, http.NoBody)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(time.Second)
	var lastReqErr error
	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), lastReqErr.Error())
		case <-ticker.C:
			res, err := http.DefaultClient.Do(request)
			if err == nil && res.StatusCode == http.StatusOK {
				_ = res.Body.Close()
				return nil
			}
			lastReqErr = err
		}
	}
}
