package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mkinney/grpc-golang-examples/src/calc/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calc(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {
	fmt.Printf("Calc function was invoked with: %v:\n", req)
	firstValue := req.GetCalc().GetFirstValue()
	secondValue := req.GetCalc().GetSecondValue()

	result := fmt.Sprintf("%d", firstValue + secondValue)
	fmt.Printf("result is: %s:\n", result)
	res := &calcpb.CalcResponse{
		Result: result,
	}
	return res, nil
}


func main() {
	fmt.Println("Starting Calc server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
  if (err != nil) {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
