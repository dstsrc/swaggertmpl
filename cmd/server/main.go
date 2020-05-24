package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/golang/glog"

	_ "github.com/dstsrc/swaggertmpl/statik"
	"github.com/rakyll/statik/fs"

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

	httpmux := http.NewServeMux()
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(statikFS)
	sh := http.StripPrefix("/swaggerui/", staticServer)
	httpmux.Handle("/swaggerui/", sh)

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = example.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	httpmux.Handle("/", mux)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Println("http start")
	return http.ListenAndServe(":8081", (httpmux))
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
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
