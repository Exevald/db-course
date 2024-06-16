package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"tree/pkg/tree/common/mysql"
	"tree/pkg/tree/infrastructure"
)

const appName = "treeOfLife"

func main() {
	ctx := context.Background()
	logger := log.New()

	conf, err := parseEnv()
	if err != nil {
		logger.Fatal(err)
	}
	err = runApp(ctx, conf, logger)
	if err != nil {
		logger.Fatal(err)
	}
}

func runApp(ctx context.Context, config *config, logger *log.Logger) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx = listenOSKillSignals(ctx)

	app := cli.App{
		Name: appName,
		Commands: []*cli.Command{
			service(config, logger),
		},
	}
	err := app.RunContext(ctx, os.Args)
	return err
}

type connectionsContainer struct {
	connector mysql.Connector
}

func newConnectionsContainer(config *config) (*connectionsContainer, error) {
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

func newDependencyContainer(config *config) (*infrastructure.DependencyContainer, error) {
	connectionsContainer, err := newConnectionsContainer(config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize connections container")
	}
	return infrastructure.NewDependencyContainer(connectionsContainer.connector)
}

func newDatabaseConnector(config *config) (mysql.Connector, error) {
	connector := mysql.NewConnector()
	err := connector.Open(config.dsn(), mysql.Config{
		MaxConnections:     config.DB.MaxConn,
		ConnectionLifetime: time.Duration(config.DB.ConnectionLifetime) * time.Second,
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
