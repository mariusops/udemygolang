package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string // lowercase field will not be exported
	Age   int    // uppercase field will be exported
}

func main() {
	u1 := user{
		First: "James", // lowercase field will not be exported
		Age:   32,      // uppercase field will be exported
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	marsUsers, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(marsUsers))

	// your code goes here
}
