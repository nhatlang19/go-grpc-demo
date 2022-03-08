package main

import (
	"context"
	"fmt"

	"github.com/nhatlang19/go-grpc-demo/calculatorpb"
	"google.golang.org/grpc"
)

var (
	server_svc = "localhost:8001"
)

func main() {
	conn, err := grpc.Dial(server_svc, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := calculatorpb.NewCalculatorServiceClient(conn)
	sum, err := client.Sum(context.Background(), &calculatorpb.SumRequest{Num1: 10, Num2: 20})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sum %v", sum)
}
