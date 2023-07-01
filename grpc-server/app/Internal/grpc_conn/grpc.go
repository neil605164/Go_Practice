package grpcconn

import (
	"go_practice/grpc-server/app/handler"
	"log"
	"net"

	"google.golang.org/grpc"
)

type IGrpcConn interface {
	GrpcConnect()
}

type GrpcConn struct {
	Health     handler.IHealth
	Calculator handler.ICalculator
}

func ProviderGrpcService(
	health handler.IHealth,
	calculator handler.ICalculator,
) IGrpcConn {
	return &GrpcConn{
		Health:     health,
		Calculator: calculator,
	}
}

func (g *GrpcConn) GrpcConnect() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := g.serviceRegister(lis)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}

func (g *GrpcConn) serviceRegister(lis net.Listener) *grpc.Server {

	grpcServer := grpc.NewServer()

	// Register Service
	g.Calculator.RegisterCalculatorService(grpcServer)
	g.Health.RegisterHealthService(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}

	return grpcServer
}
