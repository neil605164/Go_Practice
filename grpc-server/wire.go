//go:build wireinject
// +build wireinject

package main

import (
	grpcconn "go_practice/grpc-server/app/Internal/grpc_conn"

	"github.com/google/wire"
)

func Initialize() (grpcconn.IGrpcConn, error) {
	wire.Build(
		grpcconn.ProviderGrpcService,
	)

	return grpcconn.ProviderGrpcService(), nil
}
