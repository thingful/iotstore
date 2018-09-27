package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	kitlog "github.com/go-kit/kit/log"
	twrpprom "github.com/joneskoo/twirp-serverhook-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	datastore "github.com/thingful/twirp-datastore-go"

	"github.com/thingful/iotstore/pkg/rpc"
)

// Server is our top level type, contains all other components, is responsible
// for starting and stopping them in the correct order.
type Server struct {
	srv    *http.Server
	ds     *rpc.Datastore
	logger kitlog.Logger
}

// PulseHandler is the simplest possible handler function - used to expose an
// endpoint which a load balancer can ping to verify that a node is running and
// accepting connections.
func PulseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// NewServer returns a new simple HTTP server.
func NewServer(addr, connStr string, verbose bool, logger kitlog.Logger) *Server {
	ds := rpc.NewDatastore(connStr, verbose, logger)
	hooks := twrpprom.NewServerHooks(nil)

	twirpHandler := datastore.NewDatastoreServer(ds, hooks)

	// multiplex twirp handler into a mux with another handler
	mux := http.NewServeMux()
	mux.Handle(datastore.DatastorePathPrefix, twirpHandler)
	mux.HandleFunc("/pulse", PulseHandler)
	mux.Handle("/metrics", promhttp.Handler())

	// create our http.Server instance
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// return the instantiated server
	return &Server{
		srv:    srv,
		ds:     ds,
		logger: kitlog.With(logger, "module", "server"),
	}
}

// Start starts the server running. We also create a channel listening for
// interrupt signals before gracefully shutting down.
func (s *Server) Start() error {
	err := s.ds.Start()
	if err != nil {
		return err
	}

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	go func() {
		s.logger.Log("listenAddr", s.srv.Addr, "msg", "starting server")
		if err := s.srv.ListenAndServe(); err != nil {
			s.logger.Log("err", err)
			os.Exit(1)
		}
	}()

	<-stopChan
	return s.Stop()
}

// Stop stops the server running.
func (s *Server) Stop() error {
	s.logger.Log("msg", "stopping")
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	err := s.ds.Stop()
	if err != nil {
		return err
	}

	return s.srv.Shutdown(ctx)
}
