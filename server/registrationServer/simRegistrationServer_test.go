package main

import (
	"log"
	pb "sim-registration/simRegistration/registration"
	"strings"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	simRegPort = "localhost:9000"
)

func RegisterCustomerTest(t *testing.T) {
	conn, err := grpc.Dial(simRegPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	} else {
		log.Println("Connected to server")
	}
	defer conn.Close()

	client := pb.NewSimRegistrationClient(conn)

	customer := &pb.CustomerDetails{
		CustId:                 101,
		CustName:               "Vijay Upadhyay",
		Email:                  "vijay.upadhyay@nokia.com",
		AlternateContactNumber: "900000000",
		AadharNumber:           221044448888,
		Addresses: []*pb.CustomerDetails_Addrress{
			{
				HouseNumber: "34/C",
				Street:      "Gorakhnath",
				City:        "Gorakhpur",
				State:       "UP",
				Zip:         274402,
			},
			{
				HouseNumber: "36/C",
				Street:      "Friends Layout",
				City:        "Bangalore",
				State:       "Karnatak",
				Zip:         560037,
			},
		},
	}
	actualResponse, err := client.RegisterCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Error while converting from Decimal to Binary: %v", err)
	}
	if strings.Compare(actualResponse.SuccessMsg, "Success") != 0 {
		t.Errorf("Test failed, expected: '%s', got:  '%d', for input: '%d'", actualResponse.SuccessMsg, customer.CustId, customer.AadharNumber)
	}
}
