package main

import (
	"testing"
)

func TestNoValidPath(t *testing.T){
	path := "//wrong"

	if isValidPath(path) == true {
		t.Error("Expected wrong path")
	}
}

func TestValidPath(t *testing.T){
	path := "/files"

	if isValidPath(path) == false {
		t.Error("Expected correct path")
	}
}

func TestGetContentType(t *testing.T){
	if GetContentType(".html") != "text/html"{
		t.Error("Expected another")
	}
	if GetContentType(".css") != "text/css"{
		t.Error("Expected another")
	}
	if GetContentType( ".js") != "text/javascript"{
		t.Error("Expected another")
	}
	if GetContentType(".xml") != "text/xml"{
		t.Error("Expected another")
	}
	if GetContentType(".json") != "application/json"{
		t.Error("Expected another")
	}
}
