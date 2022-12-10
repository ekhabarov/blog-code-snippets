package main

import (
	"context"
	"net"
	"net/http"

	"github.com/ekhabarov/blog-code-snippets/grpc-predictable-go-interfces/repo"
	"github.com/ekhabarov/blog-code-snippets/grpc-predictable-go-interfces/service"
	"github.com/ekhabarov/blog-code-snippets/grpc-predictable-go-interfces/transport"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	lis, err := net.Listen("tcp", ":5001")
	must(err)

	s := grpc.NewServer()

	r, err := repo.New()
	must(err)

	transport.RegisterMyServiceServer(s, transport.New(service.New(r)))
	go func() {
		must(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:5001",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	must(err)

	gwmux := runtime.NewServeMux()
	must(transport.RegisterMyServiceHandler(context.Background(), gwmux, conn))

	gws := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	must(gws.ListenAndServe())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
