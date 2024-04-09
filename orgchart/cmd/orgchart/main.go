package main

import (
	"context"
	"log"
	"orgchart/pkg/orgchart/infrastructure"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
)

const appName = "orgchart"

func main() {
	ctx := context.Background()
	conf, err := parseEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = runApp(ctx, conf)

	if err != nil {
		log.Fatal(err)
	}
}

func runApp(ctx context.Context, config *config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx = listenOSKillSignals(ctx)

	app := cli.App{
		Name: appName,
		Commands: []*cli.Command{
			service(config),
		},
	}
	err := app.RunContext(ctx, os.Args)
	return err
}

func newDependencyContainer() (*infrastructure.DependencyContainer, error) {
	return infrastructure.NewDependencyContainer()
}

func listenOSKillSignals(ctx context.Context) context.Context {
	var cancelFunc context.CancelFunc
	ctx, cancelFunc = context.WithCancel(ctx)
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
		select {
		case <-ch:
			cancelFunc()
		case <-ctx.Done():
			signal.Reset()
			return
		}
	}()

	return ctx
}
