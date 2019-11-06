package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mkinney/grpc-golang-examples/src/greeting/greetingpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greeting(ctx context.Context, req *greetingpb.GreetingRequest) (*greetingpb.GreetingResponse, error) {
	fmt.Printf("Greeting function was invoked with: %v:\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetingpb.GreetingResponse{
		Result: result,
	}
	return res, nil
}


func main() {
	fmt.Println("Hello")
	gm := greetingpb.GreetingMessage{
		FirstName: "Mike",
	}
	fmt.Println(gm)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
  if (err != nil) {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetingpb.RegisterGreetingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
