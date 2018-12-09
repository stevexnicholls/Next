package api

import (
	"context"
	"log"
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
)

// Server provides an http.Server.
type Server struct {
	srv *http.Server
	rt  *next.Runtime
}

// NewServer creates and configures an APIServer serving all application routes.
func NewServer() (*Server, error) {
	log.Println("configuring server...")

	rt, err := next.NewRuntime()
	if err != nil {
		log.Fatalln(err)
	}

	k := kv.New(rt)
	b := backup.Backup{}

	// Initiate the http handler, with the objects that are implementing the business logic.
	h, err := restapi.Handler(restapi.Config{
		KvAPI:      k,
		BackupAPI:  &b,
		AuthToken:  auth.Token,
		Authorizer: nil, //auth.Request,
		Logger:     log.Printf,
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
	log.Println("starting server...")
	go func() {
		if err := srv.srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)
	// teardown logic...

	if err := srv.srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	srv.rt.Close()
	log.Println("Server gracefully stopped")
}
