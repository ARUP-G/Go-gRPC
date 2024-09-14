package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ARUP-G/Go-gRPC/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	response, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("%s", response.Message)
}
