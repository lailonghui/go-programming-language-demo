/*
@Time : 2021/4/20 18:40
@Author : lai
@Description :
@File : main
*/
package main

import (
	"context"
	"google.golang.org/grpc"
	grpc2 "lai.com/go_programming_language_demo/gRPC_Demo/demo01/client/grpc"
	"log"
	"time"
)

const (
	address = "localhost:9801"
)

func main() {
	// 连接到服务器
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := grpc2.NewHelloClient(conn)
	name := "赖龙辉22"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := client.SayHello(ctx, &grpc2.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", reply.GetMessage())
}
