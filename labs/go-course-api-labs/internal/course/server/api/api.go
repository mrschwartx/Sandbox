package api

import (
	"context"
	"net/http"

	"example.com/pkg/config"
)

var serviceName = "course-service"

type Options struct {
	config config.Config
}

type Server struct {
}

func New(opts Options) Server {
	s := Server{}

	return s
}

func (s *Server) Run(ctx context.Context) error {
	return nil
}

func (s *Server) newHTTPServer(ctx context.Context) *http.Server {
	return nil
}

func (s *Server) newGRPCServer(ctx context.Context) *http.Server {
	return nil
}

func (s *Server) healthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) readyz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
