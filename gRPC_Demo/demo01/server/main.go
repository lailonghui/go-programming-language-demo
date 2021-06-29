/*
@Time : 2021/4/20 18:40
@Author : lai
@Description :
@File : main
*/
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "lai.com/go_programming_language_demo/gRpc_Demo/demo01/server/hello"
	"log"
	"net"
)

const (
	port = ":10086"
)

// 该服务用来实现Hello服务
type server struct {
	pb.UnimplementedHelloServer
}

// SayHello implements HelloService
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "嘿嘿嘿 " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	fmt.Println("server listening on tcp:10086")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
