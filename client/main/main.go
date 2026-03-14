package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Entering client")
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("hellow world")
	}
}
