package main

import (
	"net"
	"log"
	"github.com/shettyh/threadpool"
	"bufio"
	"strings"
	"strconv"
)

type httpServer struct {
	listener net.Listener
	handlerThreadPool threadpool.ThreadPool
	handlerLoop func()
	fileReader FileReader
}

func NewHttpServer(root string, maxQueue int, maxConn int) *httpServer {
	server := &httpServer{}

	if maxQueue < 1 || maxConn < 1 {
		return nil
	}

	if thPool := threadpool.NewThreadPool(maxConn, int64(maxQueue)); thPool != nil {
		server.handlerThreadPool = *thPool
	} else {
		return nil
	}

	server.handlerLoop = func() {
		for {
			conn, err := server.listener.Accept()
			if err != nil {
				log.Fatal()
			}
			server.handlerThreadPool.Execute(NewThreadTask(func() {
				handleConnection(conn, server.fileReader)
			}))
		}
	}

	if fileReader := NewFileReader(root); fileReader != nil {
		server.fileReader = *fileReader
	} else {
		return nil
	}

	return server
}

func (server *httpServer) Start(host string, port int) error {
	var err error
	addr := host + ":" + strconv.Itoa(port)
	server.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	server.handlerLoop()
	return nil
}

const(
	Ok = "HTTP/1.1 200 OK\n\r\n\r"
	NotImplemented = "HTTP/1.1 501 Not Implemented\n\r"
	BadRequest = "HTTP/1.1 400 Bad Request\n\r"
	NotFound = "HTTP/1.1 404 Not Found\n\r"
)

func handleConnection(conn net.Conn, reader FileReader) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		requestParts := strings.Split(scanner.Text(), " ")
		if len(requestParts) < 3 {
			conn.Write([]byte(BadRequest))
			return
		}

		if requestParts[0] != "GET" {
			conn.Write([]byte(NotImplemented))
			return
		}
		if requestParts[2] != "HTTP/1.1" {
			conn.Write([]byte(BadRequest))
			return
		}

		path := requestParts[1]
		data, err := reader.ReadAllBytes(path)
		if err != nil {
			if err == PathError {
				conn.Write([]byte(BadRequest))
				return
			} else {
				conn.Write([]byte(NotFound))
				return
			}
		}

		conn.Write([]byte(Ok))
		conn.Write(data)
	}
}