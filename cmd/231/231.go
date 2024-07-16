/*
Printing and logging

fmt.Println() is used to print to the console.
log.Println() is used to print to the console and to a log file. It also adds a timestamp to the log entry.
log.Fatalln() is used to print to the console, to a log file, and to terminate the program.
log.Panicln() or panic() is used to terminate the program immediately. It runs any deferred functions before terminating the program.
You can use recover() to regain control of a panicking goroutine.
*/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, playground")
	f, err := os.Open("log.txt")
	if err != nil {
		log.Println(err)
		fmt.Println(err)
	}
	defer f.Close()

}
