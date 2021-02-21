package main

import (
	"context"
	"log"
	"net"
	"os"

	"go.seankhliao.com/testrepo-287/internal/grpcplug"
	"google.golang.org/grpc"
)

type Server struct {
	grpcplug.UnimplementedPlugServer
}

func (s *Server) SayHello(ctx context.Context, req *grpcplug.HelloReq) (*grpcplug.HelloRes, error) {
	return &grpcplug.HelloRes{Msg: "hello, " + req.N}, nil
}

func main() {
	lis, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	grpcplug.RegisterPlugServer(s, &Server{})
	s.Serve(lis)
}
