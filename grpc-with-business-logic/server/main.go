package main

import (
	"net"

	"github.com/ekhabarov/blog-code-snippets/grpc-with-business-logic/pb"
	"github.com/ekhabarov/blog-code-snippets/grpc-with-business-logic/service"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &server{
		osvc: service.New(),
	})

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		return
	}

	s.Serve(lis)
}
