package main

import (
	"bufio"
	"cards/channeldemo"
	"cards/deck"
	"cards/structsdemo"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting main....", time.Now())

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(`
		Select Programs to run:
			1. Fibonacci
			2. Cards
			3. Structs
	`)

	for scanner.Scan() {
		fmt.Println("You selected: ", scanner.Text())
		fmt.Println("*************************************************************************************************************")

		switch scanner.Text() {
		case "1":
			channeldemo.ExecuteFibonacci()
		case "2":
			deck.Execute()
		case "3":
			structsdemo.Execute()
		}

	}
}
