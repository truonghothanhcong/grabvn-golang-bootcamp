package main

import (
	pb "../feedback"
	"context"
	"log"
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
)

var client pb.FeedbackServiceClient

func main()  {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client = pb.NewFeedbackServiceClient(conn)
	router := gin.Default()

	router.POST("/feedback/add", addPassengerFeedback)
	router.POST("/feedback/passengerId", getFeedbackByPassendgerId)
	router.POST("/feedback/bookingCode", getFeedbackByBookingCode)
	router.POST("/feedback/delete", deleteFeedback)
	router.Run(":8008")
}

func addPassengerFeedback(c *gin.Context) {
	var argument *pb.FeedbackMessage

	err := c.BindJSON(&argument)
	if err != nil {
		c.String(400, "invalid param")
		return
	}

	response, err := client.AddPassengerFeedback(context.Background(), &pb.AddPassengerFeedbackRequest{FeedbackMsg: argument})
	if err != nil {
		log.Fatalf("Error when calling AddPassengerFeedback: %s", err)
	}

	c.JSON(200, response)
}

func getFeedbackByPassendgerId(c *gin.Context) {
	var argument struct {
		PassengerId int32
	}
	
	err := c.BindJSON(&argument)
	if err != nil {
		c.String(400, "invalid param %s", err)
		return
	}

	response, err := client.GetByPassengerId(context.Background(), &pb.GetByPassengerIdRequest{PassengerId: argument.PassengerId})
	if err != nil {
		log.Fatalf("Error when calling getFeedbackByPassendgerId: %s", err)
	}

	c.JSON(200, response)
}

func getFeedbackByBookingCode(c *gin.Context) {
	var argument struct {
		BookingCode string
	}

	err := c.BindJSON(&argument)
	if err != nil {
		c.String(400, "invalid param")
		return
	}

	response, err := client.GetByBookingCode(context.Background(), &pb.GetByBookingCodeRequest{BookingCode: argument.BookingCode})
	if err != nil {
		log.Fatalf("Error when calling getFeedbackByBookingCode: %s", err)
	}

	c.JSON(200, response)
}

func deleteFeedback(c *gin.Context) {
	var argument struct {
		PassengerId int32
	}

	err := c.BindJSON(&argument)
	if err != nil {
		c.String(400, "invalid param")
		return
	}

	response, err := client.DeleteByPassengerId(context.Background(), &pb.DeleteByPassengerIdRequest{PassengerId: argument.PassengerId})
	if err != nil {
		log.Fatalf("Error when calling deleteFeedback: %s", err)
	}

	c.JSON(200, response)
}