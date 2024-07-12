package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	pb "server/server/example/github.com/your_username/your_project/server/example"
	"time"
)

// grpc client

func main() {
	// grpc client
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 获取请求参数
	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
