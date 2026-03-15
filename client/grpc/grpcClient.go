package grpc

import (
	"context"
	"log"
	"time"

	pb "gauntlet/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "service"
)

var (
	addr string = "server:9000"
)

func ClientCall() {

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldnt connect %v:", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	time.Sleep(time.Second)
	r, err := c.Hello(ctx, &pb.Message{Name: "hello"})
	if err != nil {
		log.Fatalf("could not say hello :( : %v", err)
	}
	log.Printf("Hello %s", r.GetName())
}
