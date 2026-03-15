package main

import (
	"fmt"
	"gauntlet/server/grpc"
)

func main() {
	fmt.Println("starting server")
	grpc.Serve()
}
