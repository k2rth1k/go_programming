package api

import (
	gw "github.com/k2rth1k/go_programming/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
}

func Start_Server(){
	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gw.RegisterDataStructuresServer(s, &Server{})
	reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

}