package main

import (
	"net"
	"log"
)

type httpServer struct {
	listener net.Listener
	handlerLoop func()
}

func NewHttpServer() *httpServer {
	server := &httpServer{}
	server.handlerLoop = func() {
		for {
			conn, err := server.listener.Accept()
			if err != nil {
				log.Fatal()
			}
			handleConnection(conn)
		}
	}
	return server
}

func (server *httpServer) Start() error {
	var err error
	addr := "127.0.0.1:3000"
	server.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	server.handlerLoop()
	return nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("HTTP/1.1 200 OK\n\r\n\r"))
}
