package main

import (
	"context"
	"flag"
	"github.com/k2rth1k/go_programming/pkg/api"
	algo "github.com/k2rth1k/go_programming/pkg/proto/algorithms"
	ds "github.com/k2rth1k/go_programming/pkg/proto/data-structures"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
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
	err := ds.RegisterDataStructuresHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	err=algo.RegisterAlgorithmsHandlerFromEndpoint(ctx,mux,*grpcServerEndpoint,opts)
	if err!=nil{
		return err
	}
	go api.StartServer()
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