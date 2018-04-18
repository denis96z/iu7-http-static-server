package main

import (
	"net"
	"log"
	"github.com/shettyh/threadpool"
	"bufio"
	"strings"
)

type httpServer struct {
	listener net.Listener
	handlerThreadPool threadpool.ThreadPool
	handlerLoop func()
}

func NewHttpServer() *httpServer {
	server := &httpServer{}
	server.handlerThreadPool = *threadpool.NewThreadPool(10, 1000)
	server.handlerLoop = func() {
		for {
			conn, err := server.listener.Accept()
			if err != nil {
				log.Fatal()
			}
			server.handlerThreadPool.Execute(NewThreadTask(func() {
				handleConnection(conn)
			}))
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

const(
	NotImplemented = "HTTP/1.1 501 Not Implemented\n\r"
	BadRequest = "HTTP/1.1 400 Bad Request\n\r"
	Ok = "HTTP/1.1 200 OK\n\r\n\r"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		requestParts := strings.Split(scanner.Text(), " ")
		if  requestParts[0] != "GET" {
			conn.Write([]byte(NotImplemented))
			return
		}

		if requestParts[2] != "HTTP/1.1" {
			conn.Write([]byte(BadRequest))
			return
		}

		conn.Write([]byte(Ok))
	}
}