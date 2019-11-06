package main

import (
	"fmt"
	"log"

	"github.com/mkinney/greet-service-go/src/greeting/greetingpb"
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
	fmt.Printf("created client: %f", c)

}
