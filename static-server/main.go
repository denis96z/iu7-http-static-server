package main

import "fmt"

func main() {
	if server := NewHttpServer(); server != nil {
		if server.Start() != nil {
			fmt.Println("failed to start server...")
		}
	} else {
		fmt.Println("invalid configuration...")
	}
}