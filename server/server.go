package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nhatlang19/go-grpc-demo/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("Sum called")
	resp := &calculatorpb.SumResponse{
		Total: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("Server starting ....")
	err = s.Serve(listen)

	if err != nil {
		panic(err)
	}

}
