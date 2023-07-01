package grpcconn

import (
	"go_practice/grpc-server/app/handler"
	"go_practice/pkg/pb/message"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type IGrpcConn interface {
	GrpcConnect()
}

type grpcConn struct{}

func ProviderGrpcService() IGrpcConn {
	return &grpcConn{}
}

func (g *grpcConn) GrpcConnect() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := g.serviceRegister(lis)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}

func (g *grpcConn) serviceRegister(lis net.Listener) *grpc.Server {

	grpcServer := grpc.NewServer()

	// Register Calculator Service
	message.RegisterCalculatorServiceServer(grpcServer, &handler.Server{})

	// Register Health Service
	healthServer := health.NewServer()
	healthServer.SetServingStatus("grpc.health.v1.Health", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}

	return grpcServer
}
