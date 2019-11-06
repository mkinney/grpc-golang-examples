package main

import (
	"fmt"
	"context"
	"log"

	"github.com/mkinney/grpc-golang-examples/src/greeting/greetingpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello")
	gm := greetingpb.GreetingMessage{
		FirstName: "Mike",
	}
	fmt.Println(gm)

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer cc.Close()

	c := greetingpb.NewGreetingServiceClient(cc)
	//fmt.Printf("created client: %f", c)

	doUnary(c)
}

func doUnary(c greetingpb.GreetingServiceClient) {
	fmt.Println("Starting to do Unary RPC")
	req := &greetingpb.GreetingRequest{
		Greeting: &greetingpb.GreetingMessage{
			FirstName: "Mike",
			LastName: "Kinney",
		},
	}

	res, err := c.Greeting(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greeting RPC: %v:", err)
	}
	log.Printf("response from Greeting: %v", res.Result)

}
