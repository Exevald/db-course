package server

import (
	"sync/atomic"

	"github.com/pkg/errors"
)

type ServeFunc func() error

type StopFunc func() error

const (
	serverIsCreated int32 = iota
	serverIsRunning
	serverIsStopped
)

type server struct {
	serveFunc ServeFunc
	stopFunc  StopFunc
	state     int32
}

func newServer(serve ServeFunc, stop StopFunc) *server {
	return &server{
		serveFunc: serve,
		stopFunc:  stop,
		state:     serverIsCreated,
	}
}

func (s *server) serve() error {
	if !atomic.CompareAndSwapInt32(&s.state, serverIsCreated, serverIsRunning) {
		if atomic.LoadInt32(&s.state) == serverIsRunning {
			return errAlreadyRun
		}
		return errTryRunStoppedServer
	}
	return s.serveFunc()
}

func (s *server) stop() error {
	stopped := atomic.CompareAndSwapInt32(&s.state, serverIsCreated, hubIsStopped) ||
		atomic.CompareAndSwapInt32(&s.state, serverIsRunning, serverIsStopped)

	if !stopped {
		return errAlreadyStopped
	}
	return s.stopFunc()
}

var errAlreadyStopped = errors.New("server is not running, can't change server state")
var errAlreadyRun = errors.New("server is running, can't change server state to running")
var errTryRunStoppedServer = errors.New("server is stopped, can't change server state to running")
