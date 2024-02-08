package main

import (
	"fmt"
)

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	doors  int
	color  string
}

func printVehicle(v vehicle) {
	engineType := "electric"
	if !v.engine.electric {
		engineType = "gas"
	}
	fmt.Println(v.make, v.model, v.doors, v.color, engineType)
}

func main() {

	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Tesla",
		model: "S",
		doors: 4,
		color: "red",
	}

	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Tesla",
		model: "X",
		doors: 4,
		color: "white",
	}

	printVehicle(v1)
	printVehicle(v2)
}
