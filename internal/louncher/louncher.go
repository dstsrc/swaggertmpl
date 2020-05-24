package louncher

import (
	"context"
	"log"
	"net"
	"net/http"

	example "github.com/dstsrc/swaggertmpl/pkg/pb"
	_ "github.com/dstsrc/swaggertmpl/statik"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
)

type louncher struct {
	server   example.YourServiceServer
	endpoint string
}

func New(server example.YourServiceServer, endpoint string) *louncher {
	return &louncher{server: server, endpoint: endpoint}
}

func (l *louncher) Run() {
	go func() {
		if err := l.runHTTP(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := l.runGRPC(); err != nil {
		log.Fatal(err)
	}
}

func (l *louncher) runGRPC() error {
	s := grpc.NewServer()
	example.RegisterYourServiceServer(s, l.server)

	lis, err := net.Listen("tcp", l.endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("grpc start")
	return s.Serve(lis)
}

func (l *louncher) runHTTP() error {
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
	err = example.RegisterYourServiceHandlerFromEndpoint(ctx, mux, l.endpoint, opts)
	if err != nil {
		return err
	}

	httpmux.Handle("/", mux)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Println("http start")
	return http.ListenAndServe(":8081", (httpmux))
}
