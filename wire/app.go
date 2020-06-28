package main

import (
	"net"
	"wire/rpci"

	"google.golang.org/grpc"
)

// App contains minimal list of dependencies to be able to start an application.
type App struct {
	// listener is a TCP listener which is used by gRPC server.
	listener net.Listener
	// gRPC serer itself.
	gsrv *grpc.Server
	// gRPC server implementation. It's not used here directly, but it must be
	// initialized for registering. gRPC server.
	rpcImpl *rpci.Server
}

// Start start gRPC server.
func (a App) Start() error {
	return a.gsrv.Serve(a.listener)
}
