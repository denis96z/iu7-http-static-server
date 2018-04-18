package main

import (
	"testing"
	"reflect"
)

func TestNoConfigFile(t *testing.T){
	path := "wrong.json"
	var config ServerConfig
	if config.FromFile(path) == nil {
		t.Error("Expected file not found")
	}
}

func TestConfigFile(t *testing.T){
	path := "config.json"
	var config ServerConfig
	if config.FromFile(path) != nil {
		t.Error("Error opening file")
	}

	expectedConfig := ServerConfig {
		Host:  "127.0.0.1",
		Port: 3000,
		RootDir: ".",
		MaxConn: 3000,
		MaxQueueLen: 100,
		MaxFileSize: 1024,
	}

	if !reflect.DeepEqual(expectedConfig, config) {
		t.Error("Error read file")
	}
}
