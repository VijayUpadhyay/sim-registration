package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"fmt"
	duplicateCustomerCheck "sim-registration/checkDuplicateCustomer/custValidator"
	pb "sim-registration/simRegistration/registration"
)

const (
	port        = ":9000"
	chkCustPort = "localhost:50058"
)

type server struct {
	saveCustomerDetails []*pb.CustomerDetails
	pb.UnimplementedSimRegistrationServer
}

func (s *server) RegisterCustomer(ctx context.Context, in *pb.CustomerDetails) (*pb.CreateResponseMessage, error) {
	// Check whether the customer is registered or not
	conn, err := grpc.Dial(chkCustPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server for checking whether customer is already registered or not: %v", err)
	} else {
		log.Print("Connected to server")
	}
	defer conn.Close()

	c := duplicateCustomerCheck.NewCheckDuplicateCustomerClient(conn)

	val, err := c.CheckDuplicateCustomerBeforeCreate(context.Background(), &duplicateCustomerCheck.RequestMsg{AadharNumber: in.AadharNumber})
	if err != nil {
		log.Fatalf("Failed to communicate: %s", err)
	}

	if !val.CustExist {
		fmt.Printf("Welcome %v\n", in.CustName)
	} else {
		fmt.Printf("Hi %v!\n", in.CustName)
		//displayCustDetails(in.AadharNumber) -- to fetch customer according to their Aadhar Number
	}
	//
	s.saveCustomerDetails = append(s.saveCustomerDetails, in)
	return &pb.CreateResponseMessage{CustId: in.CustId, SuccessMsg: "Success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Print("Listening to port ", port)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterSimRegistrationServer(s, &server{})
	//s.Serve(lis)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve")
	}
}
