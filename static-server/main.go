package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("config file path expected...")
		return
	}
	var config ServerConfig
	if config.FromFile("./config.json") != nil {
		fmt.Println("failed to parse config file...")
		return
	}
	if server := NewHttpServer(-1, config.RootDir,
			config.MaxQueueLen, config.MaxConn);
			server != nil {
		if server.Start(config.Host, config.Port) != nil {
			fmt.Println("failed to start server...")
		}
	} else {
		fmt.Println("invalid configuration...")
	}
}