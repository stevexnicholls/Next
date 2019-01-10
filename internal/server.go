package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/spf13/viper"
	"github.com/stevexnicholls/next/auth"
	"github.com/stevexnicholls/next/internal/backup"
	"github.com/stevexnicholls/next/internal/kv"
	"github.com/stevexnicholls/next/internal/runtime"
	"github.com/stevexnicholls/next/restapi"
	log "github.com/stevexnicholls/next/logger"
)

// Server provides an http.Server.
type Server struct {
	srv *http.Server
	rt  *next.Runtime
}

// NewServer creates and configures an APIServer serving all application routes.
func NewServer() (*Server, error) {
	log.Info("configuring server...")

	rt, err := next.NewRuntime()
	if err != nil {
		log.Fatal(err)
	}

	k := kv.New(rt)
	b := backup.Backup{}

	// Initiate the http handler, with the objects that are implementing the business logic.
	h, err := restapi.Handler(restapi.Config{
		KvAPI:      k,
		BackupAPI:  &b,
		AuthToken:  auth.Token,
		Authorizer: auth.Request,
		Logger:     log.Infof,
	})
	if err != nil {
		log.Fatal(err)
	}

	var addr string
	port := viper.GetString("port")

	// allow port to be set as localhost:3000 in env during development to avoid "accept incoming network connection" request on restarts
	if strings.Contains(port, ":") {
		addr = port
	} else {
		addr = ":" + port
	}

	srv := http.Server{
		Addr:    addr,
		Handler: h,
	}

	return &Server{srv: &srv, rt: rt}, nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *Server) Start() {
	log.Info("starting server...")

	protocol := viper.GetString("protocol")
	cert := viper.GetString("tls_cert")
	key := viper.GetString("tls_key")

	go func() {
		switch protocol {
		case "http":
			if err := srv.srv.ListenAndServe(); err != http.ErrServerClosed {
				panic(err)
			}
		case "https":
			if err := srv.srv.ListenAndServeTLS(cert, key); err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()
	log.Infof("listening on %s://%s", protocol, srv.srv.Addr)
	
	if viper.Get("log_level") == "debug" {
		log.Infof("keystore path: %s", viper.Get("store_path"))
		log.Infof("keystore bucket: %s", viper.Get("store_bucket"))
		log.Infof("api key: %s", viper.Get("api_key"))
		log.Infof("log path: %s", viper.Get("log_path"))	
	}
	
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Infof("shutting down server... reason: %s", sig)
	// teardown logic...

	if err := srv.srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	srv.rt.Close()
	log.Info("server gracefully stopped")
}
