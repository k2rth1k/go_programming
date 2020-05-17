package api

import (
	"context"
	ds "github.com/k2rth1k/go_programming/pkg/proto/data-structures"
	"io"
	"log"
	"sort"
	"time"
)


func(s *Server) RearrangeArrayAlternately(ctx context.Context,req *ds.RearrangeArrayAlternatelyRequest)  (*ds.RearrangeArrayAlternatelyResponse,error){
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
	solution:=&ds.RearrangeArrayAlternatelyResponse{ArrangedArray: result,Array:req.Array}
	return solution,nil
}

func(s *Server) StreamClientSide(srv ds.DataStructures_StreamClientSideServer) error {
	log.Println("[called StreamClientSide rpc]")
	result := ""
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return srv.SendAndClose(&ds.StreamClientSideResponse{Sol: result})
		}
		if err != nil {
			log.Fatalf("[failed to recieve req:%v...]", err)
		}
		result = result + req.Message + "!\n"
	}
	return nil
}

func(s *Server) ServerSideStream(req *ds.ServerSideStreamRequest,srv ds.DataStructures_ServerSideStreamServer) error{
	log.Println("[called ServerSideStream rpc]")
	for i:=int64(0);i<=req.Count;i++{
		err:=srv.Send(&ds.ServerSideStreamResponse{Number: i})
		if err != nil {
			log.Fatalf("[failed to send res:%v...]", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func(s *Server) BidirectionalStreaming(srv ds.DataStructures_BidirectionalStreamingServer) error{
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

func stream(count int64,srv ds.DataStructures_BidirectionalStreamingServer) error{
	for i:=int64(0);i<=count;i++{
		err:=srv.Send(&ds.BidirectionalStreamResponse{Numbers: i})
		if err != nil {
			log.Fatalf("[failed to send res:%v...]", err)
			return err
		}
		time.Sleep(300 * time.Millisecond)
	}
	return nil
}