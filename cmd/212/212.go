package main

import (
	"fmt"
	"runtime"
)

func main() {
	// print os architecture
	fmt.Println(runtime.GOARCH)
	// print os operating system
	fmt.Println(runtime.GOOS)
}
