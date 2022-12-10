proto:
	 protoc \
		 -I ./transport \
		 -I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		 --go_out ./transport \
		 --go_opt paths=source_relative \
		 --go-grpc_out ./transport \
		 --go-grpc_opt paths=source_relative \
		 --grpc-gateway_out ./transport \
		 --grpc-gateway_opt paths=source_relative \
		 ./transport/transport.proto

