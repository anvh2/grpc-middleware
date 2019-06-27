package greeter

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-middleware/grpc-gen/proto"

	"google.golang.org/grpc"
)

//Server ...
type Server struct{}

//NewServer ...
func NewServer() *Server {
	return &Server{}
}

//Run ...
func (s *Server) Run() error {
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return err
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	pb.RegisterGreeterServiceServer(grpcServer, s)

	fmt.Println("Server is runing on port 8000")

	return grpcServer.Serve(lis)
}

//Say ...
func (s *Server) Say(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	msg := "Hello " + req.GetName()

	return &pb.Response{
		Message: msg,
	}, nil
}
