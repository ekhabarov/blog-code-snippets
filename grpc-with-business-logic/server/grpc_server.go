package main

import (
	"context"
	"log"
	"time"

	"github.com/ekhabarov/blog-code-snippets/grpc-with-business-logic/pb"
	"github.com/ekhabarov/blog-code-snippets/grpc-with-business-logic/service"
)

// server implements gRPC Server OrderServiceServer interface from
// pb/order.pb.go.
type server struct {
	osvc service.OrderSvc
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("GET request: %#v\n", req)
	o, err := s.osvc.Get(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	items := []*pb.Item{}

	for _, i := range o.Items {
		items = append(items, &pb.Item{
			Id:   int64(i.ID),
			Name: i.Name,
			Qty:  int32(i.Qty),
		})
	}

	order := &pb.Order{
		Id:        int64(o.ID),
		Number:    o.Num,
		Status:    o.Status,
		CreatedAt: time.Now().Unix(),
		Items:     items,
	}

	return &pb.GetResponse{
		Order: order,
	}, nil
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("ADD request: %#v\n", req)
	o := req.Body.Order

	id, err := s.osvc.Add(ctx, &service.Order{
		ID:      int(o.Id),
		Num:     o.Number,
		Status:  o.Status,
		Items:   nil, // skip it again
		Comment: "new order",
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddResponse{
		Id:     int64(id),
		Status: "added",
	}, nil
}
