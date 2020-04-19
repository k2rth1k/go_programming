package main
import (
	"context"
	"flag"
	"github.com/k2rth1k/go_programming/pkg/api"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "github.com/k2rth1k/go_programming/pkg/proto"  // Update
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:9090", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterDataStructuresHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	go startServer()
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func startServer(){
	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gw.RegisterDataStructuresServer(s, &api.Server{})
	reflection.Register(s)
	go func(s *grpc.Server,lis net.Listener) {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}(s,lis)
}