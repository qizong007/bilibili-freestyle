package main

import (
	"io"
	"os"
	"strings"
)

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func WriteStringsToFile(list []string, filePath string) error {
	writeString := strings.Join(list, "\n")
	var f *os.File
	var err error
	if checkFileIsExist(filePath) {
		f, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	} else {
		f, err = os.Create(filePath)
	}
	if f == nil {
		return err
	}
	defer f.Close()
	_, err = io.WriteString(f, writeString)
	return err
}
