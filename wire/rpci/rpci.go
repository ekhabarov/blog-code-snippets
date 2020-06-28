package rpci

import (
	"context"
	"wire/pb"
	"wire/service"

	"google.golang.org/grpc"
)

type Server struct {
	xlogic service.Xlogic
}

func New(x service.Xlogic, gsrv *grpc.Server) *Server {
	s := &Server{xlogic: x}
	pb.RegisterEntityRPCServer(gsrv, s)

	return s
}

func (s *Server) EntityByID(ctx context.Context, req *pb.EntityByIDRequest) (*pb.EntityByIDResponse, error) {
	xe, err := s.xlogic.GetEntity(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.EntityByIDResponse{Entity: &pb.Entity{
		Id:   int64(xe.ID),
		Name: xe.Name,
	}}, nil
}
