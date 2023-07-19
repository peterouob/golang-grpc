package main

import (
	"google.golang.org/grpc"
	pb "grpc/proto"
	"log"
	"net"
)

const (
	port = ":8083"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("Faild to listen on port server: %v", err)
	}
	log.Printf("Server listening on port [%v]", port)
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("Failed to start grpc server: %v", err)
	}

}
