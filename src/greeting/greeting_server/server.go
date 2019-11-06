package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mkinney/greet-service-go/src/greeting/greetingpb"
	"google.golang.org/grpc"
)

type server struct{}

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
