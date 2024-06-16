package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"tree/api/server/treepublic"
	"tree/pkg/tree/common/server"
	"tree/pkg/tree/infrastructure"
	"tree/pkg/tree/infrastructure/transport/common"
	"tree/pkg/tree/infrastructure/transport/publicapi"
)

func service(config *config, logger *log.Logger) *cli.Command {
	return &cli.Command{
		Name:  "service",
		Usage: "Runs application as http server",
		Action: func(c *cli.Context) error {
			dependencyContainer, err := newDependencyContainer(config)
			if err != nil {
				return errors.Wrap(err, "could not create dependency container")
			}
			return runService(config, dependencyContainer, logger)
		},
	}
}

func runService(
	config *config,
	dependencyContainer *infrastructure.DependencyContainer,
	logger *log.Logger,
) error {
	api := publicapi.NewPublicAPI()
	errChan := make(chan struct{})
	serverHub := server.NewHub(errChan)

	serveHTTP(config, serverHub, api, logger)
	return serverHub.Wait()
}

func serveHTTP(
	config *config,
	serverHub *server.Hub,
	api publicapi.PublicAPI,
	logger *log.Logger,
) {
	_, cancel := context.WithCancel(context.Background())
	var httpServer *http.Server
	serverHub.Serve(func() error {
		publicAPIHandler := treepublic.NewStrictHandler(api, []treepublic.StrictMiddlewareFunc{
			common.NewLoggingMiddleware(logger),
			publicapi.NewErrorsMiddleware(),
		})
		router := mux.NewRouter()
		router.HandleFunc("/resilience/ready", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		}).Methods(http.MethodGet)
		router.PathPrefix("/api/v1/treeoflife").Handler(treepublic.Handler(publicAPIHandler))
		httpServer = &http.Server{
			Addr:              config.Service.ServeRESTAddress,
			ReadHeaderTimeout: 10 * time.Second,
			ReadTimeout:       time.Hour,
			WriteTimeout:      time.Hour,
		}
		return httpServer.ListenAndServe()
	}, func() error {
		cancel()
		return httpServer.Shutdown(context.Background())
	})
}
