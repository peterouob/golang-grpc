package main

import (
	pb "grpc/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStream(stream pb.GreetService_SayHelloClientStreamServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got Request with name: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}
