package main

import (
	"fmt"

	"github.com/mariusops/udemygolang/cmd/244/pkg/dog"
)

func main() {
	dogYears := dog.Years(5)
	fmt.Println(dogYears)
}
