package api

import (
	"context"
	pb "github.com/k2rth1k/go_programming/pkg/proto"
	"io"
	"log"
	"sort"
	"time"
)


func(svc *Server) RearrangeArrayAlternately(ctx context.Context,req *pb.RearrangeArrayAlternatelyRequest)  (*pb.RearrangeArrayAlternatelyResponse,error){
	log.Println("[called RearrangeArrayAlternately rpc]")
	var arr []int64
	for _,number:=range req.Array{
		arr=append(arr,number)
	}
	sortedArr:=arr
	sort.Slice(sortedArr, func(i, j int) bool {
		return sortedArr[i]< sortedArr[j]
	})
	var result []int64
	length:=len(sortedArr)
	for i:=0;i<=(length-1)/2;i++{
		result = append(result, sortedArr[length-i-1])
		if length-1-i!=i{
			result = append(result, sortedArr[i])
		}
	}
	solution:=&pb.RearrangeArrayAlternatelyResponse{ArrangedArray:result,Array:req.Array}
	return solution,nil
}

func(svc *Server) StreamClientSide(srv pb.DataStructures_StreamClientSideServer) error {
	log.Println("[called StreamClientSide rpc]")
	result := ""
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return srv.SendAndClose(&pb.StreamClientSideResponse{Sol: result})
		}
		if err != nil {
			log.Fatalf("[failed to recieve req:%v...]", err)
		}
		result = result + req.Message + "!\n"
	}
	return nil
}

func(svc *Server) ServerSideStream(req *pb.ServerSideStreamRequest,srv pb.DataStructures_ServerSideStreamServer) error{
	log.Println("[called ServerSideStream rpc]")
	for i:=int64(0);i<=req.Count;i++{
		err:=srv.Send(&pb.ServerSideStreamResponse{Number:i})
		if err != nil {
			log.Fatalf("[failed to send res:%v...]", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func(svc *Server) BidirectionalStreaming(srv pb.DataStructures_BidirectionalStreamingServer) error{
	log.Println("[called BidirectionalStream rpc]")
	for {
		req, err := srv.Recv()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatalf("[failed to recieve req:%v...]", err)
		}
		err=stream(req.Count,srv)
		if err != nil {
			return err
		}
	}
	return nil
}

func stream(count int64,srv pb.DataStructures_BidirectionalStreamingServer) error{
	for i:=int64(0);i<=count;i++{
		err:=srv.Send(&pb.BidirectionalStreamResponse{Numbers: i})
		if err != nil {
			log.Fatalf("[failed to send res:%v...]", err)
			return err
		}
		time.Sleep(300 * time.Millisecond)
	}
	return nil
}