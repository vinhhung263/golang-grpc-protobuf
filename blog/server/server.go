package main

import pb "github.com/vinhhung263/grpc-go-course/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
