package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func getAbsolutePath(filename string) (string, error) {
	absulutePath, err := filepath.Abs(filename)
	return absulutePath, err
}

func readFile(file *os.File) []byte {
	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	return data
}

func main() {
	a, b := 15, 13
	if a > b {
		fmt.Println("a é maior do que b")
	} else if a < b {
		fmt.Println("a é menor do que b")
	} else {
		fmt.Println("a e b são iguais")
	}

	filename := "hello.txt"
	absoluteFilePath, pathErr := getAbsolutePath(filename)
	if pathErr != nil {
		log.Panic(pathErr)
	}

	file, err := os.Open(absoluteFilePath)
	if err != nil {
		log.Panic(err)
	}

	data := readFile(file)
	fmt.Println(string(data))
}
