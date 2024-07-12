package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "server/server/example/github.com/your_username/your_project/server/example"
)

func main() {
	fmt.Println("server begin***")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &pb.UnimplementedHelloServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
