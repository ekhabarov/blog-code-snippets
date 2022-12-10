package transport

import (
	"context"

	"github.com/ekhabarov/blog-code-snippets/grpc-predictable-go-interfces/paginator"
	"github.com/ekhabarov/blog-code-snippets/grpc-predictable-go-interfces/service"
)

type Server struct {
	svc service.Service
}

func New(svc service.Service) *Server {
	return &Server{
		svc: svc,
	}
}

func (s *Server) List(ctx context.Context, r *Request) (*Response, error) {
	entities, err := s.svc.List(int(r.Page), int(r.Limit))
	if err != nil {
		return nil, err
	}

	var out []*Entity

	for _, e := range entities {
		out = append(out, &Entity{Id: int32(e.ID), Name: e.Name})
	}

	return &Response{Entities: out}, nil
}

func (s *Server) ListWithApplier(ctx context.Context, r *Request) (*Response, error) {
	entities, err := s.svc.ListWithApplier(paginator.FromRequest(r))
	if err != nil {
		return nil, err
	}

	var out []*Entity

	for _, e := range entities {
		out = append(out, &Entity{Id: int32(e.ID), Name: e.Name})
	}

	return &Response{Entities: out}, nil
}

func (s *Server) mustEmbedUnimplementedMyServiceServer() {
	panic("not implemented") // TODO: Implement
}
