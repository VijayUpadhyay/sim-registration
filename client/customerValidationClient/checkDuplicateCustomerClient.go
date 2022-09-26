package main

import (
	"log"
	pb "sim-registration/checkDuplicateCustomer/custvalidator"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const chkCustPortClient = "localhost:50058"

func main() {
	conn, err := grpc.Dial(chkCustPortClient, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	} else {
		log.Println("Connected to server")
	}
	defer conn.Close()

	c := pb.NewCheckDuplicateCustomerClient(conn)

	val, err := c.CheckDuplicateCustomerBeforeCreate(context.Background(), &pb.RequestMsg{AadharNumber: 222233334444})
	if err != nil {
		log.Fatalf("Failed to communicate: %s", err)
	}
	log.Println("Registered Customer: ", val.CustExist)
}
