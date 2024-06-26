package mapsdemo

import (
	"fmt"
)

func Execute() {
	colors := map[string]string{
		"red": "#ff0000",
	}
	colors["white"] = "#fff"

	fmt.Println(colors)
	fmt.Println(len(colors))
	delete(colors, "red")

	if _, ok := colors["red"]; ok {
		fmt.Println("Red color exists")
	} else {
		fmt.Println("Red color does not exist")
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
