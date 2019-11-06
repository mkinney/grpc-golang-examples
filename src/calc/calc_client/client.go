package main

import (
	"fmt"
	"context"
	"log"

	"github.com/mkinney/grpc-golang-examples/src/calc/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello")

	cm := calcpb.Calc {
		FirstValue: 10,
		SecondValue: 3,
	}
	fmt.Println(cm)

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer cc.Close()

	c := calcpb.NewCalcServiceClient(cc)
	//fmt.Printf("created client: %f", c)

	doUnary(c)
}

func doUnary(c calcpb.CalcServiceClient) {
	fmt.Println("Starting to do Unary RPC")
	req := &calcpb.CalcRequest{
		Calc: &calcpb.Calc{
			FirstValue: 10,
			SecondValue: 3,
		},
	}

	res, err := c.Calc(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Calc RPC: %v:", err)
	}
	log.Printf("response from Calc: %v", res.Result)

}
