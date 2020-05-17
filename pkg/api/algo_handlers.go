package api

import (
	"context"
	algo "github.com/k2rth1k/go_programming/pkg/proto/algorithms"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func(s *Server) FindMinimumNumber(ctx context.Context,req *algo.FindMinimumNumberRequest) (*algo.FindMinimumNumberResponse,error){
	log.Println("[called FindMinimumNumber rpc]")
	arr:=req.Numbers
	length:=len(req.Numbers)
	if length>0 {
		min := arr[0]
		for index := 1; index < length; index++ {
			if arr[index] < min {
				min = arr[index]
			}
		}
		return &algo.FindMinimumNumberResponse{MinimumNumber: min}, nil
	}else {
		return nil,status.Error(codes.InvalidArgument,"empty array is sent")
	}
}
