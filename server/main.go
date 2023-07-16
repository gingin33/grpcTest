package main

import (
	"fmt"
	"net"

	"github.com/jon20/grpc-stream-sample/server/grpc"
	"google.golang.org/grpc"
)

func main() {
	port := "8080"

	//データを待つ
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening at " + port + "...")

	//gRPCサーバを立てる
	server := grpc.NewServer()

	handler.NewUploadServer(server)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
