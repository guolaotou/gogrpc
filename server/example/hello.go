package example

// import pb
import (
	pb "server/server/example/github.com/your_username/your_project/server/example"
)

// Server rpc服务端
type Server struct {
	pb.UnimplementedHelloServiceServer
}
