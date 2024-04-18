// writer interface
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, playground")

	/*
		Fprintln and WriteString are two functions that write to the writer interface
		They both got the method signature:
			func (w Writer) Write(p []byte) (n int, err error)
		We're using os.Stdout as the writer
		os.Stdout is a writer that writes to the standard output
		os.Stdout writes to a file descriptor 1 which is the standard output
	*/

	// Fprintln writes to the writer and adds a newline character at the end
	fmt.Fprintln(os.Stdout, "Hello, playground")

	// WriteString writes to the writer without adding a newline character
	io.WriteString(os.Stdout, "Hello, playground")

}
