package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"log"
)

func TestServerOk(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := "GET /config.json HTTP/1.1"
		reader := NewFileReader(".")
		if reader == nil {
			t.Error("failed to create reader")
		}
		code, _, body := handleRequest(req, *reader)
		w.WriteHeader(code)
		if code == OkCode {
			fmt.Fprintln(w, body)
		}
	}))

	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("expected 200 OK")
	}
}

func TestServerFileNotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := "GET /wrong.json HTTP/1.1"
		reader := NewFileReader(".")
		if reader == nil {
			t.Error("failed to create reader")
		}
		code, _, body := handleRequest(req, *reader)
		w.WriteHeader(code)
		if code == OkCode {
			fmt.Fprintln(w, body)
		}
	}))

	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != NotFoundCode {
		t.Error("expected 404 Not Found")
	}
}

func TestServerBadRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := "GET /config.json HTP/1.1"
		reader := NewFileReader(".")
		if reader == nil {
			t.Error("failed to create reader")
		}
		code, _, body := handleRequest(req, *reader)
		w.WriteHeader(code)
		if code == OkCode {
			fmt.Fprintln(w, body)
		}
	}))

	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != BadRequestCode {
		t.Error("expected 400 Bad Request")
	}
}

func TestServerNotImplementedMethod(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := "POST /wrong.json HTTP/1.1"
		reader := NewFileReader(".")
		if reader == nil {
			t.Error("failed to create reader")
		}
		code, _, body := handleRequest(req, *reader)
		w.WriteHeader(code)
		if code == OkCode {
			fmt.Fprintln(w, body)
		}
	}))

	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != NotImplementedCode{
		t.Error("expected 501 Not Implemented")
	}
}