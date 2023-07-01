//go:build wireinject
// +build wireinject

package main

import (
	grpcconn "go_practice/grpc-server/app/Internal/grpc_conn"
	"go_practice/grpc-server/app/handler"

	"github.com/google/wire"
)

func Initialize() (grpcconn.GrpcConn, error) {
	panic(wire.Build(
		handler.ProviderHealthCli,
		handler.ProviderCalculatorCli,

		wire.Struct(new(grpcconn.GrpcConn), "*"),
	))
}
