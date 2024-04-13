package main

import (
	"context"
	"github.com/gorilla/mux"
	stdio "io"
	"net/http"
	"orgchart/pkg/orgchart/infrastructure"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func service(config *config) *cli.Command {
	return &cli.Command{
		Name:  "service",
		Usage: "Runs application as http server",
		Action: func(c *cli.Context) error {
			dependencyContainer, err := newDependencyContainer()
			if err != nil {
				return errors.Wrap(err, "could not create dependency container")
			}
			return runService(c.Context, config, dependencyContainer)
		},
	}
}

func runService(
	ctx context.Context,
	config *config,
	dependencyContainer *infrastructure.DependencyContainer,
) error {
	return serveHTTP(config)
}

func serveHTTP(
	config *config,
) error {
	var httpServer *http.Server

	//publicAPIHandler := orgchartpublic.NewStrictHandler(api, []orgchartpublic.StrictMiddlewareFunc{})
	router := mux.NewRouter()
	router.HandleFunc("/resilience/ready", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = stdio.WriteString(w, http.StatusText(http.StatusOK))
	}).Methods(http.MethodGet)
	//
	//router.PathPrefix("/api/v1/orgchart").Handler(orgchartpublic.Handler(publicAPIHandler))
	//
	httpServer = &http.Server{
		Handler:           router,
		Addr:              config.Service.ServeRESTAddress,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       time.Hour,
		WriteTimeout:      time.Hour,
	}

	return httpServer.ListenAndServe()
}
