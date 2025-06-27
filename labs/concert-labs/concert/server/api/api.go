package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"example.com/concert/server/booking"
	"example.com/concert/server/catalog"
	"example.com/internal/config"
	"example.com/internal/gateway"
	v1 "example.com/pkg/api/concert/v1"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serviceName = "concert-service"

type Opts struct {
	config config.Server
}

type Server struct {
	opts Opts
}

func NewServer(opts Opts) Server {
	log.Debug().
		Str("postgres", fmt.Sprintf("%s:%s/%s", opts.config.DB.Host, opts.config.DB.Port, opts.config.DB.Name)).
		Str("redis", opts.config.Redis.GetAddr()).
		Msg("server configuration loaded successfully")

	return Server{
		opts: opts,
	}
}

func (s *Server) Run(ctx context.Context) error {
	log.Info().Str("service", serviceName).Msg("initializing service...")

	grpcServer := s.newGRPCServer(ctx)

	go func() {
		addr := s.opts.config.GRPC.GetAddr()
		log.Info().Str("address", addr).Msg("starting gRPC server")

		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal().Err(err).Str("address", addr).Msg("unable to bind to gRPC address")
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("gRPC server encountered a fatal error")
		}
	}()

	httpServer := s.newHTTPServer(ctx)

	go func() {
		addr := s.opts.config.HTTP.GetAddr()
		log.Info().Str("address", addr).Msg("starting HTTP server for gRPC-Gateway and API documentation")

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("HTTP server encountered a fatal error")
		}
	}()

	<-ctx.Done()
	log.Warn().Msg("shutdown signal received; initiating graceful shutdown sequence")

	gracefulShutdownPeriod := 30 * time.Second
	shutdownCtx, cancel := context.WithTimeout(context.Background(), gracefulShutdownPeriod)
	defer cancel()

	log.Info().Msg("shutting down HTTP server")
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("HTTP server shutdown failed")
	} else {
		log.Info().Msg("HTTP server shut down gracefully")
	}

	log.Info().Msg("shutting down gRPC server")
	grpcServer.GracefulStop()
	log.Info().Msg("gRPC server shut down gracefully")

	log.Info().Msg("performing final resource cleanup (TODO)")
	return nil
}

func (s *Server) newGRPCServer(ctx context.Context) *grpc.Server {
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	bookingServer := booking.NewServer()
	catalogServer := catalog.NewServer()

	v1.RegisterBookingServiceServer(grpcServer, bookingServer)
	v1.RegisterCatalogServiceServer(grpcServer, catalogServer)

	return grpcServer
}

func (s *Server) newHTTPServer(ctx context.Context) *http.Server {
	conn, err := grpc.NewClient(
		s.opts.config.GRPC.GetAddr(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to dial grpc server: %v", err)
	}

	gwmux := runtime.NewServeMux()
	gateway.MustRegisterGWHandler(ctx, v1.RegisterCatalogServiceHandler, gwmux, conn)
	gateway.MustRegisterGWHandler(ctx, v1.RegisterBookingServiceHandler, gwmux, conn)

	mux := mux.NewRouter()
	mux.HandleFunc("/healthz", s.healthz())
	mux.HandleFunc("/readyz", s.readyz())
	mux.PathPrefix("/debug/").Handler(http.DefaultServeMux)

	api := mux.PathPrefix("/api/concert").Subrouter()
	api.Use()
	api.PathPrefix("/v1").Handler(gwmux)

	swaggerUI := http.StripPrefix("/swagger/", http.FileServer(http.Dir("./third_party/OpenAPI/")))
	mux.PathPrefix("/swagger/").Handler(swaggerUI)

	gwServer := &http.Server{
		Addr:    s.opts.config.HTTP.GetAddr(),
		Handler: mux,
	}
	return gwServer
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
