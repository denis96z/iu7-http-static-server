package main

import (
	"encoding/json"
	"io/ioutil"
)

type ServerConfig struct {
	Host string 	`json:"Host"`
	Port int 		`json:"Port"`
	RootDir string	`json:"RootDirectory"`
	MaxConn int		`json:"MaxConnections"`
	MaxQueueLen int `json:"MaxQueue"`
	MaxFileSize int `json:"MaxSizeFile"`
}

func (config *ServerConfig) FromFile(path string) error {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(buffer, config)
}