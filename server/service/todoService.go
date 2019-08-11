package service

import (
	pb "../../proto"
	"../dbworker"

	"context"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Server struct {
	ToDoDBWorker dbworker.ToDo
}

func (s *Server) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	id, _ := uuid.NewV4()
	req.Item.Id = id.String()
	err := s.ToDoDBWorker.Insert(req.Item)

	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}

	return &pb.CreateTodoResponse{Item: req.Item}, nil
}
func (s *Server) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	item, err := s.ToDoDBWorker.Get(req.Id)
	if err != nil {
		return nil, grpc.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
	}

	return &pb.GetTodoResponse{Item: item}, nil
}
func (s *Server) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoResponse, error) {
	var items []*pb.Todo
	items, err := s.ToDoDBWorker.List(req.Limit, req.NotCompleted)
	if err != nil {
		return nil, grpc.Errorf(codes.NotFound, "Could not list items from the database: %s", err)
	}

	return &pb.ListTodoResponse{Items: items}, nil
}
func (s *Server) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	err := s.ToDoDBWorker.Delete(req.Id)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not delete item from the database: %s", err)
	}
	return &pb.DeleteTodoResponse{}, nil
}
