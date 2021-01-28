/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-28 10:31:29
 * @LastEditors: Caoxiang
 */
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/newinternetboy/grpc_examples/proto"
	"google.golang.org/grpc"
)

const PORT = "8520"

//Search服务
type SearchService struct{}

//Search服务中对应的具体Search方法
func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

//Search服务中对应的具体SearchDotor方法
func (s *SearchService) SearchDoctor(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: "Search Doctor"}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	fmt.Println("net.Listen running...")
	server.Serve(lis)
}
