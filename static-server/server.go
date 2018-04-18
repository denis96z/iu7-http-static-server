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

func NewHttpServer(numRequests int, root string, maxQueue int, maxConn int) *httpServer {
	server := &httpServer{}

	if maxQueue < 1 || maxConn < 1 {
		return nil
	}

	if thPool := threadpool.NewThreadPool(maxConn, int64(maxQueue)); thPool != nil {
		server.handlerThreadPool = *thPool
	} else {
		return nil
	}

	loopFunc := func () {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Fatal()
		}
		server.handlerThreadPool.Execute(NewThreadTask(func() {
			handleConnection(conn, server.fileReader)
		}))
	}

	server.handlerLoop = func() {
		if numRequests < 0 {
			for {
				loopFunc()
			}
		} else {
			for i := 0; i < numRequests; i++ {
				loopFunc()
			}
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
	OkCode = 200
	Ok = "HTTP/1.1 200 OK\r\n"

	NotImplementedCode = 501
	NotImplemented = "HTTP/1.1 501 Not Implemented\r\n"

	BadRequestCode = 400
	BadRequest = "HTTP/1.1 400 Bad Request\r\n"

	NotFoundCode = 404
	NotFound = "HTTP/1.1 404 Not Found\r\n"
)

func handleConnection(conn net.Conn, reader FileReader) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		code, headers, body := handleRequest(scanner.Text(), reader)
		switch code {
		case OkCode:
			conn.Write([]byte(Ok))
			for _, h := range headers {
				conn.Write([]byte(h + "\r\n"))
			}
			conn.Write([]byte("\r\n"))
			conn.Write(body)
		case BadRequestCode:
			conn.Write([]byte(BadRequest))
		case NotFoundCode:
			conn.Write([]byte(NotFound))
		case NotImplementedCode:
			conn.Write([]byte(NotImplemented))
		}
	}
}

func handleRequest(req string, reader FileReader) (int, []string, []byte) {
	requestParts := strings.Split(req, " ")
	if len(requestParts) < 3 {
		return BadRequestCode, nil, nil
	}

	if requestParts[0] != "GET" {
		return NotImplementedCode, nil, nil
	}
	if requestParts[2] != "HTTP/1.1" {
		return BadRequestCode, nil, nil
	}

	path := requestParts[1]
	data, err := reader.ReadAllBytes(path)
	if err != nil {
		if err == PathError {
			return BadRequestCode, nil, nil
		} else {
			return NotFoundCode, nil, nil
		}
	}

	var headers []string
	headers = append(headers, "Connection: close")
	headers = append(headers, "Server: iu7-http-static-server")
	if contType := GetContentType(path); contType != "" {
		headers = append(headers, "Content-Type: " + contType + "")
	}
	headers = append(headers, "Content-Length: " + strconv.Itoa(len(data)))

	return OkCode, headers, data
}