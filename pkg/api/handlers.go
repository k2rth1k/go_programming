package api

import (
	"context"
	pb "github.com/k2rth1k/go_programming/pkg/proto"
	"sort"
)


func(svc *Server) RearrangeArrayAlternately(ctx context.Context,req *pb.RearrangeArrayAlternatelyRequest)  (*pb.RearrangeArrayAlternatelyResponse,error){
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