/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-28 11:46:17
 * @LastEditors: Caoxiang
 */
package main

import (
	"context"
	"log"

	pb "github.com/newinternetboy/grpc_examples/proto"
	"google.golang.org/grpc"
)

const PORT = "8520"

func main() {
	// Search()
	SearchDoctor()
}

func Search() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err:%v", err)
	}
	log.Printf("resp: %s", resp.GetResponse())
}

func SearchDoctor() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	resp, err := client.SearchDoctor(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err:%v", err)
	}
	log.Printf("resp: %s", resp.GetResponse())

}
