package main

import (
	"log"

	pb "github.com/ARUP-G/Go-gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	// Make a client
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Arup", "Ron", "Con"},
	}

	callSayHello(client)
	callSayHelloServerStream(client, names)
}
