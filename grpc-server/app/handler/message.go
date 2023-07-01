package handler

import (
	"context"
	"fmt"
	"go_practice/pkg/pb/message"

	"google.golang.org/grpc"
)

type ICalculator interface {
	RegisterCalculatorService(grpcServer *grpc.Server)
}

type calculator struct {
	message.UnsafeCalculatorServiceServer
}

func ProviderCalculatorCli() ICalculator {
	return &calculator{}
}

func (c *calculator) RegisterCalculatorService(grpcServer *grpc.Server) {
	message.RegisterCalculatorServiceServer(grpcServer, &calculator{})
}

func (c *calculator) Sum(ctx context.Context, req *message.CalculatorRequest) (*message.CalculatorResponse, error) {
	fmt.Printf("Sum function is invoked with %v \n", req)

	a := req.GetA()
	b := req.GetB()

	fmt.Println("req.A=====>", req.A)
	fmt.Println("req.B=====>", req.B)

	res := &message.CalculatorResponse{
		Result: a + b,
	}

	return res, nil
}
