package gateway

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type registerFunc func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

// MustRegisterGWHandler is a convenience function to register a gateway handler.
func MustRegisterGWHandler(ctx context.Context, register registerFunc, mux *runtime.ServeMux, conn *grpc.ClientConn) {
	err := register(ctx, mux, conn)
	if err != nil {
		panic(err)
	}
}
