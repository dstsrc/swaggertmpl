package main

import (
	"flag"

	"github.com/dstsrc/swaggertmpl/internal/louncher"

	"github.com/dstsrc/swaggertmpl/internal/app"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func main() {
	flag.Parse()

	a := app.NewApp()

	l := louncher.New(a, *grpcServerEndpoint)

	l.Run()

}
