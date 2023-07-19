package main

import (
	"context"
	pb "grpc/proto"
)

// 照著proto裡面的rpc寫

// 這邊因為返回所以是響應而非流
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "hello",
	}, nil
}
