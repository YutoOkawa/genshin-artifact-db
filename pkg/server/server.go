package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Server          http.Server
	Port            string
	ShutdownTimeout int
}

func NewServer(port string, handler http.Handler, shutdownTimeout int) *Server {
	return &Server{
		Server: http.Server{
			Addr:    port,
			Handler: handler,
		},
		Port:            port,
		ShutdownTimeout: shutdownTimeout,
	}
}

func (s *Server) Start() chan error {
	errorCh := make(chan error, 1)
	go func() {
		if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
			errorCh <- err
		}
	}()
	return errorCh
}

func (s *Server) Shutdown() {
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.ShutdownTimeout)*time.Second)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
