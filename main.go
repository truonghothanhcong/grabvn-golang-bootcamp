package main

import (
	pb "./ssh"
	"context"
	"net"
	"log"
	"google.golang.org/grpc"
)

type server struct {}

func (*server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{
		ResponseCode: 0,
		UserResponse: &pb.User{
			Id: 1,
			Name: "Hien Xau Trai",
		},
	}, nil
}

func (*server) GetListUser(ctx context.Context, req *pb.GetListUserRequest) (*pb.GetListUserResponse, error) {
	return nil, nil
}
func (*server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, nil
}
func (*server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, nil
}
func (*server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}

func main()  {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}