package main

import (
	pb "./route"
	"context"
	"net"
	"log"
	"google.golang.org/grpc"
)

type server struct {
	
}

func (*server) SayHelloToTheWorld(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{
		Msg:"ahihi",
	}, nil
}


func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	pb.RegisterMyRouteServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}