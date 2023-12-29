package main

import (
	"log"
	"os"
	"testing"
)

const (
	fileName string = "testdata/test.txt"
)

func readFile() string {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		log.Print("error while reading file", err)
		return ""
	}
	return string(fileData)
}

func TestWC(t *testing.T) {
	type wcFunc func(string) int
	tests := []struct {
		wcFunc   wcFunc
		expected int
	}{
		{CountBytes, 342190},
		{CountLines, 7145},
		{CountWords, 58164},
		{CountChars, 339292},
	}
	data := readFile()
	for _, test := range tests {
		ret := test.wcFunc(data)
		if ret != test.expected {
			t.Errorf("expected %d, got %d", test.expected, ret)
		}
	}
}
