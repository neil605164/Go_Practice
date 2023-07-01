package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("starting gRPC server...")

	grpc, err := Initialize()
	if err != nil {
		log.Fatal(err)
	}
	grpc.GrpcConnect()
}
