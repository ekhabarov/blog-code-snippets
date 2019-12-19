package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ekhabarov/blog-code-snippets/simple-grpc-service/pb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("127.0.0.1:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewCalcClient(cc)
	resp, err := client.Add(context.Background(), &pb.Request{A: 2, B: 2})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("response is: %d\n", resp.C)
}
