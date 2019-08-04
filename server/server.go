package main

import (
	pb "../feedback"
	"context"
	"net"
	"log"
	"google.golang.org/grpc"
)

var ResponseSuccess = &pb.ResponseCode {
	Code: 0,
	Msg: "Success",
}

var ResponseNotFound = &pb.ResponseCode {
	Code: 1,
	Msg: "Not Found",
}

var ResponseExist = &pb.ResponseCode {
	Code: 2,
	Msg: "Exist",
}

type server struct {
	Feedbacks map[string]*pb.FeedbackMessage
}

func (s *server) AddPassengerFeedback(ctx context.Context, req *pb.AddPassengerFeedbackRequest) (*pb.AddPassengerFeedbackResponse, error) {
	var Code = ResponseSuccess
	var ResponseData = req.FeedbackMsg

	if _, ok := s.Feedbacks[req.FeedbackMsg.BookingCode]; ok {
		Code = ResponseExist
	} else {
		s.Feedbacks[req.FeedbackMsg.BookingCode] = req.FeedbackMsg
	}

	return &pb.AddPassengerFeedbackResponse{
		ResCode: Code,
		FeedbackMsg: ResponseData,
	}, nil
}

func (s *server) GetByPassengerId(ctx context.Context, req *pb.GetByPassengerIdRequest) (*pb.GetByPassengerIdResponse, error) {
	var ResponseCode = ResponseSuccess
	var ResponseDatas []*pb.FeedbackMessage

	for _, v := range s.Feedbacks {
		if v.PassengerId == req.PassengerId {
			ResponseDatas = append(ResponseDatas, v)
		}
	}

	if len(ResponseDatas) == 0 {
		ResponseCode = ResponseNotFound
	}

	return &pb.GetByPassengerIdResponse{
		ResCode: ResponseCode,
		FeedbackMsgs: ResponseDatas,
	}, nil
}

func (s *server) GetByBookingCode(ctx context.Context, req *pb.GetByBookingCodeRequest) (*pb.GetByBookingCodeResponse, error) {
	var Code = ResponseNotFound
	var ResponseData *pb.FeedbackMessage

	if v, ok := s.Feedbacks[req.BookingCode]; ok {
		Code = ResponseSuccess
		ResponseData = v
	}

	return &pb.GetByBookingCodeResponse{
		ResCode: Code,
		FeedbackMsg: ResponseData,
	}, nil
}

func (s *server) DeleteByPassengerId(ctx context.Context, req *pb.DeleteByPassengerIdRequest) (*pb.DeleteByPassengerIdResponse, error) {
	var ResponseCode = ResponseSuccess
	var BookingCodes []string

	for k, v := range s.Feedbacks {
		if v.PassengerId == req.PassengerId {
			BookingCodes = append(BookingCodes, k)
		}
	}

	NumberDeleted := len(BookingCodes)
	for _, k := range BookingCodes {
		delete(s.Feedbacks, k)
	}

	return &pb.DeleteByPassengerIdResponse{
		ResCode: ResponseCode,
		NumberOfDeleted: int32(NumberDeleted),
	}, nil
}


func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	pb.RegisterFeedbackServiceServer(s, &server{
		Feedbacks: make(map[string]*pb.FeedbackMessage),
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}