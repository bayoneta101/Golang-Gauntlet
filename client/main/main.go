package main

import (
	"fmt"
	grpc "gauntlet/client/grpc"
)

func main() {
	fmt.Println("Entering client")
	grpc.ClientCall()
}
