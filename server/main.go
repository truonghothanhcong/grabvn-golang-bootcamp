package main

import (
	pb "../proto"
	"./service"
	"./dbworker"

	"fmt"
	"net"
	"log"
	"time"
	"google.golang.org/grpc"
	"github.com/go-pg/pg"
)


func main() {
	grpcAddress := "localhost:5001"
	// httpAddress := "localhost:5002"

	fmt.Printf("Connecting into DB\n")
	// Connect to PostgresQL
	db := pg.Connect(&pg.Options{
		User:                  "postgres",
		Password:              "example",
		Database:              "todo",
		Addr:                  "localhost" + ":" + "5433",
		RetryStatementTimeout: true,
		MaxRetries:            4,
		MinRetryBackoff:       250 * time.Millisecond,
	})
	// Create Table from Todo struct generated by gRPC
	db.CreateTable(&pb.Todo{}, nil)
	
	s := grpc.NewServer()
	todoDBWorker := dbworker.ToDoImpl{DB: db}
	pb.RegisterTodoServiceServer(s, &service.Server{
		ToDoDBWorker: todoDBWorker,
	})

	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("can not listen tcp grpcAddress %s: %v", grpcAddress, err)
	}

	fmt.Printf("Serving GRPC at %s.\n", grpcAddress)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
