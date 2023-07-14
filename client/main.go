package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/jon20/grpc-stream-sample/proto"

	"google.golang.org/grpc"
)

func main() {
	connect, _ := grpc.Dial("10.0.2.62:8080", grpc.WithInsecure())

	defer connect.Close()
	uploadhalder := upload.NewUploadHandlerClient(connect)
	stream, err := uploadhalder.Upload(context.Background())
	err = Upload(stream)
	if err != nil {
		fmt.Println(err)
	}
}

func Upload(stream upload.UploadHandler_UploadClient) error {
	saveDockerImage()
	imageName := "tmp.tar"

	err := os.MkdirAll("images", 0777)
	if err != nil {
		return err
	}

	file, _ := os.Open("./images/" + imageName)
	defer file.Close()
	buf := make([]byte, 1024)
	for {

		_, err := file.Read(buf)

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		stream.Send(&upload.UploadRequest{VideoData: buf})
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Println(resp.UploadStatus)

	return nil
}

func saveDockerImage() error {
	flag.Parse()

	args := flag.Args()
	err := exec.Command("docker", "save", args[0], "-o", "./images/tmp.tar").Run()
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
