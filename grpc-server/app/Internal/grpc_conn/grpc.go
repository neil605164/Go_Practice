package grpcconn

import (
	"go_practice/grpc-server/app/handler"
	"log"
	"net"

	"google.golang.org/grpc"
)

type IGrpcConn interface {
	Run()
}

type GrpcConn struct {
	Health     *handler.HealthCli
	Calculator *handler.Calculator
}

func ProviderGrpcService(
	health *handler.HealthCli,
	calculator *handler.Calculator,
) IGrpcConn {
	return &GrpcConn{
		Health:     health,
		Calculator: calculator,
	}
}

func (g *GrpcConn) Run() {
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
