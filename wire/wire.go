// +build wireinject

package main

import (
	"database/sql"
	"net"
	"wire/repo"
	"wire/rpci"
	"wire/service"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

func NewListener() (net.Listener, error) {
	return net.Listen("tcp4", "0.0.0.0:5000")
}

func NewGRPCServer() *grpc.Server {
	return grpc.NewServer()
}

func DBConn() (*sql.DB, error) {
	return sql.Open("mysql", "127.0.0.1:3306")
}

func initApp() (*App, error) {
	wire.Build(
		rpci.New,
		NewListener,
		NewGRPCServer,
		repo.Provider,
		DBConn,
		service.Provider,
		wire.Struct(new(App), "*"),
	)

	return &App{}, nil
}
