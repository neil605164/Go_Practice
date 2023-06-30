package handler

import (
	"context"
	"fmt"
	"go_practice/pkg/pb/message"
)

type Server struct {
	message.UnsafeCalculatorServiceServer
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
