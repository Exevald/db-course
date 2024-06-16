package main

import (
	"context"
	"fmt"
	"net/http"
	"orgchart/api/server/orgchartpublic"
	"orgchart/pkg/common/mysql"
	"orgchart/pkg/common/server"
	"orgchart/pkg/integrationaltests/infrastructure"
	"orgchart/pkg/integrationaltests/tests"
	serviceinfrastruture "orgchart/pkg/orgchart/infrastructure"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

const (
	waitTimeout = 30 * time.Second
)

// default:"orgchart-tests-db-dev:3306"

type Config struct {
	DBHost               string `envconfig:"db_host" default:"orgchart-tests-db-dev:3306"`
	DBName               string `envconfig:"db_name" default:"orgchart-test"`
	DBUser               string `envconfig:"db_user" default:"orgchart-test"`
	DBPassword           string `envconfig:"db_password" default:"LxilKD9Pbe"`
	DBMaxConn            int    `envconfig:"db_max_conn" default:"0"`
	DBConnectionLifetime int    `envconfig:"db_conn_lifetime" default:"0"`

	ServiceHTTPEndpoint string `envconfig:"service_http_endpoint" default:"http://orgchart-dev-for-tests:8082"`
	ServeRESTAddress    string `envconfig:"serve_rest_address" default:":8082"`
}

func (c *Config) dsn() mysql.DSN {
	return mysql.DSN{
		Host:     c.DBHost,
		Database: c.DBName,
		User:     c.DBUser,
		Password: c.DBPassword,
	}
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
	orgchartClient, err := orgchartpublic.NewClientWithResponses(config.ServiceHTTPEndpoint)
	if err != nil {
		log.WithError(err).Fatal("failed to connect to public API")
	}
	publicAPI := infrastructure.NewOrgchartPublicAPI(orgchartClient)
	dependencyContainer, err := newDependencyContainer(&config)
	if err != nil {
		log.WithError(err).Fatal("could not create dependency container")
	}

	log.Info("starting tests...")

	tests.RunTests(publicAPI, dependencyContainer.BranchService(), dependencyContainer.EmployeeService())

	log.Info("all tests completed successfully")
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

func newDependencyContainer(config *Config) (*serviceinfrastruture.DependencyContainer, error) {
	connectionsContainer, err := newConnectionsContainer(config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize connections container")
	}
	return serviceinfrastruture.NewDependencyContainer(connectionsContainer.connector)
}

type connectionsContainer struct {
	connector mysql.Connector
}

func newConnectionsContainer(config *Config) (*connectionsContainer, error) {
	container := &connectionsContainer{}
	containerBuilder := func() error {
		connector, err := newDatabaseConnector(config)
		if err != nil {
			return err
		}
		container.connector = connector
		return nil
	}
	return container, containerBuilder()
}

func newDatabaseConnector(config *Config) (mysql.Connector, error) {
	connector := mysql.NewConnector()
	fmt.Println(config.DBHost)
	err := connector.Open(config.dsn(), mysql.Config{
		MaxConnections:     config.DBMaxConn,
		ConnectionLifetime: time.Duration(config.DBConnectionLifetime) * time.Second,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed database connection")
	}
	err = connector.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "failed database ping")
	}
	return connector, nil
}
