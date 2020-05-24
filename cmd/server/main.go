package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/dstsrc/swaggertmpl/internal/app"

	example "github.com/dstsrc/swaggertmpl/pkg/pb"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := example.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Println("http start")
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()

	a := app.NewApp()

	s := grpc.NewServer()
	example.RegisterYourServiceServer(s, a)

	lis, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err := run(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("grpc start")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
