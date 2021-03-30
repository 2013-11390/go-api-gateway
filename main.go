package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

const (
	HTTPPortNumber = ":9000"
	gRPCPortNumber = ":9001"
)

func ServeHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	return http.ListenAndServe(HTTPPortNumber, mux)
}

func ServeGRPC() error {
	lis, err := net.Listen("tcp", gRPCPortNumber)

	if err != nil {
		return err
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	lis, err := net.Listen("tcp", portNumber)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	grpcServer.Serve(lis)
}
