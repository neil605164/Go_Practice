package handler

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type IHealth interface {
	RegisterHealthService(grpcServer *grpc.Server)
}

type healthCli struct {
}

func ProviderHealthCli() IHealth {
	return &healthCli{}
}

func (h *healthCli) RegisterHealthService(grpcServer *grpc.Server) {

	// Register Health Service
	healthServer := health.NewServer()
	healthServer.SetServingStatus("grpc.health.v1.Health", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(grpcServer, healthServer)
}
