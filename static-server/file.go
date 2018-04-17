package main

import (
	"io/ioutil"
	"errors"
	"unicode"
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
	cPrev := '\n'
	for _, c := range path {
		if cPrev == '/' && c == '/' {
			return false
		}
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) &&
			c != '/' && c != '.' {
			return false
		}
		cPrev = c
	}
	return true
}