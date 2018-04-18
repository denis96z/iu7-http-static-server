package main

import "fmt"

func main() {
	var config ServerConfig
	if config.FromFile("./config.json") != nil {
		fmt.Println("failed to parse config file...")
		return
	}
	if server := NewHttpServer(); server != nil {
		if server.Start() != nil {
			fmt.Println("failed to start server...")
		}
	} else {
		fmt.Println("invalid configuration...")
	}
}