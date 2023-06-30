package main

import (
	"context"
	"fmt"
	"go_practice/pkg/proto/message"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	message.UnsafeCalculatorServiceServer
}

func main() {
	fmt.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	message.RegisterCalculatorServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}

func (c *Server) Sum(ctx context.Context, req *message.CalculatorRequest) (*message.CalculatorResponse, error) {
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

func (c *Server) GetFibonacci(req *message.GetFibonacciRequest, stream message.CalculatorService_GetFibonacciServer) error {
	position := req.GetNum()
	cache := make([]int64, position+1)
	result := fibMemo(position, cache)

	for _, num := range result {
		stream.Send(&message.GetFibonacciResponse{
			Num: int64(num),
		})
		time.Sleep(1 * time.Second)
	}

	return nil
}

func fibMemo(position int64, cache []int64) []int64 {
	if cache[position] != 0 {
		return cache
	} else {
		if position <= 2 {
			cache[position] = 1
		} else {
			cache[position] = fibMemo(position-1, cache)[position-1] + fibMemo(position-2, cache)[position-2]
		}

		return cache
	}
}
