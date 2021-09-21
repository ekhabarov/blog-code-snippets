package server

import (
	"context"

	"github.com/ekhabarov/blog-code-snippets/bazel-k8s-envoy/service-one/pb"
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Body: "Hello, " + req.Name,
	}, nil
}
