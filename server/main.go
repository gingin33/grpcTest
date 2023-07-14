package main

import (
	"net"

	"github.com/jon20/grpc-stream-sample/server/grpc"
	"google.golang.org/grpc"
)

func main() {
	//データを待つ
	lis, err := net.Listen("tcp", "10.0.2.117:8080")
	if err != nil {
		panic(err)
	}

	//gRPCサーバを立てる
	server := grpc.NewServer()

	handler.NewUploadServer(server)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
