package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ekhabarov/blog-code-snippets/grpc-with-business-logic/pb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("127.0.0.1:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewOrderServiceClient(cc)

	r, err := client.Get(context.Background(), &pb.GetRequest{Id: 2})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("order:\nID: %d, Number: %s, Status: %s\n", r.Order.Id, r.Order.Number, r.Order.Status)

	for _, i := range r.Order.Items {
		fmt.Printf("\titem: ID: %d, Name: %s, Qty: %d\n", i.Id, i.Name, i.Qty)
	}

	newOrder, err := client.Add(context.Background(), &pb.AddRequest{
		Body: &pb.AddRequest_Body{
			Order: &pb.Order{
				Id:        1,
				Number:    "123",
				Status:    "new",
				CreatedAt: time.Now().Unix(),
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nnew order ID: %d, Status: %s\n", newOrder.Id, newOrder.Status)
}
