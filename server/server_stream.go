package main

import (
	"log"
	"time"

	pb "github.com/ARUP-G/Go-gRPC/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		// Adding delay to simulate long running process
		time.Sleep(2 * time.Second)
	}
	return nil
}
