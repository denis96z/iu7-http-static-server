package main

import (
	"io/ioutil"
	"errors"
	"unicode"
	"path/filepath"
)

// Менеджер чтения файла.
type FileReader struct {
	baseRoot string
}

var (
	// Ошибка в имени файла.
	PathError = errors.New("invalid path")
)

// Создает менеджер чтения файла.
func NewFileReader(root string) *FileReader {
	return &FileReader{root}
}

// Считывает данные из файла в буфер.
func (reader *FileReader) ReadAllBytes(path string) ([]byte, error) {
	fullPath := reader.baseRoot + path
	if isValidPath(fullPath) {
		return ioutil.ReadFile(fullPath)
	}
	return nil, PathError
}

// Проверяет корретность пути в файловой системе.
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

// Возвращает тип контента в зависимости от
// расширения файла. Для неизвестного расширения
// возвращает пустую строку.
func GetContentType(path string) string {
	ct := ""
	switch filepath.Ext(path) {
	case ".html":
		ct = "text/html"
	case ".css":
		ct = "text/css"
	case ".js":
		ct = "text/javascript"
	case ".xml":
		ct = "text/xml"
	case ".json":
		ct = "application/json"
	}
	return ct
}