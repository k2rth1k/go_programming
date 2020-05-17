package api

import (
	algo "github.com/k2rth1k/go_programming/pkg/proto/algorithms"
	ds "github.com/k2rth1k/go_programming/pkg/proto/data-structures"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
}

type Algorithms struct{

}

func StartServer(){
	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ds.RegisterDataStructuresServer(s, &Server{})
	algo.RegisterAlgorithmsServer(s,&Server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}