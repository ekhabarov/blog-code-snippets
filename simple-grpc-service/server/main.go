package main

import (
	"context"
	"fmt"
	"net"

	"github.com/ekhabarov/blog-code-snippets/simple-grpc-service/pb"
	grpc "google.golang.org/grpc"
)

// server implements gRPC Server CalcServer interface from pb/calc.pb.go
type server struct{}

func (s *server) Add(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	fmt.Printf("got request: A = %d, B = %d\n", req.A, req.B)
	return &pb.Response{
		C: req.A + req.B,
	}, nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterCalcServer(s, &server{})

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		return
	}

	s.Serve(lis)
}
