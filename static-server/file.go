package main

import (
	"io/ioutil"
	"errors"
)

type FileReader struct {
	baseRoot string
}

var (
	PathError = errors.New("invalid path")
)

func NewFileReader(root string) *FileReader {
	return &FileReader{root}
}

func (reader *FileReader) ReadAllBytes(path string) ([]byte, error) {
	fullPath := reader.baseRoot + path
	if isValidPath(fullPath) {
		return ioutil.ReadFile(fullPath)
	}
	return nil, PathError
}

func isValidPath(path string) bool {
	//TODO
	return true
}