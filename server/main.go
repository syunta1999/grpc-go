package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	go_protocol_buffer "grpc-go-practice/go-protocol-buffer"
)

type server struct {
	go_protocol_buffer.UnimplementedPinPonServiceServer
}

func (s *server) Send(ctx context.Context, req *go_protocol_buffer.PinPonRequest) (*go_protocol_buffer.PinPonResponse, error) {

	resWords := ""

	if req.Words == "Pin" {
		resWords = "Pon!"
	} else {
		resWords = "Please need words 'Pin'!"
	}

	res := &go_protocol_buffer.PinPonResponse{
		Words: resWords,
	}

	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	creds := insecure.NewCredentials()
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	go_protocol_buffer.RegisterPinPonServiceServer(grpcServer, &server{})

	fmt.Println("server is runnig...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
