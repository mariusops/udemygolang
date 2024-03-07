// wrapper function to read file

// what is a wrapper function?
// a function that calls another function and returns its result

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("Error reading file in main: %s", err)
	}
	fmt.Println(string(xb))
}

func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %s", err)
	}
	return xb, nil
}
