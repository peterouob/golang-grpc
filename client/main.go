package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/proto"
	"log"
)

const (
	port = ":8083"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	name := &pb.NameList{
		Names: []string{"peter", "Alice", "Bob"},
	}
	//callHello(client)
	//callHelloServerStream(client, name)
	//callHelloClientStream(client, name)
	callHelloBidirectionalStream(client, name)
}
