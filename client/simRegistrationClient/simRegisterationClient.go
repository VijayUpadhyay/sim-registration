package main

import (
	"log"

	pb "sim-registration/simRegistration/registration"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:9000"
)

func RegisterCustomer(client pb.SimRegistrationClient, customerDetails *pb.CustomerDetails) {
	resp, err := client.RegisterCustomer(context.Background(), customerDetails)
	if err != nil {
		log.Fatalf("Could not register Customer: %v", err)
	}
	if resp.SuccessMsg == "Success" {
		log.Printf("A new Customer has been registered with id: %d", resp.CustId)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
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

	// Create a new customer
	RegisterCustomer(client, customer)

	customer = &pb.CustomerDetails{
		CustId:                 102,
		CustName:               "Ram",
		Email:                  "ram@nokia.com",
		AlternateContactNumber: "1232234354",
		AadharNumber:           222244448888,
		Addresses: []*pb.CustomerDetails_Addrress{
			{
				HouseNumber: "44/C",
				Street:      "Chinnapannahalli",
				City:        "Bangalore",
				State:       "Karnataka",
				Zip:         560068,
			},
		},
	}

	// Create a new customer
	RegisterCustomer(client, customer)
}
