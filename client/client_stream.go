package main

import (
	"context"
	pb "grpc/proto"
	"log"
	"time"
)

func callHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("client stream started")
	stream, err := client.SayHelloClientStream(context.Background())
	if err != nil {
		log.Fatalf("could not send name :%v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("could not send stream :%v", err)
		}
		log.Printf("send the request with name:%v", name)
		time.Sleep(2 * time.Second)
	}
	//接收響應
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not close stream recv :%v", err)
	}
	log.Printf("%v", res.Messages)
}
