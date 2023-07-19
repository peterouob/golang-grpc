package main

import (
	"context"
	pb "grpc/proto"
	"io"
	"log"
)

func callHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("streaming started")
	stream, err := client.SayHelloServerStream(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send name :%v", err)
	}
	//無論收到什麼流，客戶端像服務器發出請求，服務器都會返回一個流，客戶端現在要處理
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error with stream: %v", err)
		}
		log.Println(message)
	}
	log.Printf("stream finished")
}
