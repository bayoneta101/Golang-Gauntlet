package grpc

import (
	"context"
	"fmt"

	pb "gauntlet/gen"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Hello(_ context.Context, in *pb.Message) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Reply{Name: "response " + in.GetName()}, nil
}

var port = "9000"

func Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
