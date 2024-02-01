package main

import (
	"fmt"
	"io"
	"os"
)

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	return file
}

func splitIntoWords(contents string) map[string]int {
	const space = ' ' // Change space from string to rune
	var word string
	m := make(map[string]int)

	for _, char := range contents {
		if char == '\n' || char == '\t' || char == '\r' {
			continue
		}
		if char != space {
			word += string(char)
		} else {
			m[word]++
			word = ""
		}
	}
	return m
}

func main() {
	// open file named great-gatsby.txt
	// read contents of file into a string

	file := openFile("great-gatsby.txt")

	if file != nil {
		fmt.Println("File opened successfully")
	} else {
		fmt.Println("File not opened")
	}

	contents, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// print file contents
	//fmt.Println(string(contents))

	// split file contents into a slice of words
	words := splitIntoWords(string(contents))

	// print words
	for word, count := range words {
		fmt.Println(word, count)
	}

	// look for the word with the most occurrences
	// print the word and the number of occurrences
	var maxWord string
	var maxCount int
	for word, count := range words {
		if count > maxCount {
			maxCount = count
			maxWord = word
		}
	}
	fmt.Println("The word with the most occurrences is", maxWord, "with", maxCount, "occurrences")

}
