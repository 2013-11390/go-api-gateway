package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	listenAddress = ":9001"
)

func main() {
	lis, err := net.Listen("tcp", listenAddress)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	grpcServer.Serve(lis)
}
