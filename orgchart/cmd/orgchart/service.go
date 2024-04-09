package main

import (
	"context"
	stdio "io"
	"net/http"

	"orgchart/pkg/orgchart/infrastructure"

	"github.com/gorilla/mux"
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
	return nil
}

func serveHTTP(
	config *config,
) {
	router := mux.NewRouter()

	router.HandleFunc("/resilience/ready", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = stdio.WriteString(w, http.StatusText(http.StatusOK))
	}).Methods(http.MethodGet)

	//router.PathPrefix("/").Handler(func(w http.ResponseWriter, _ *http.Request) {})
}
