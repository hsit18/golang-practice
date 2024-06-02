package main

import (
	"bufio"
	"cards/channeldemo"
	"cards/deck"
	"cards/goroutinedemo"
	"cards/grpcdemo"
	"cards/mapsdemo"
	"cards/notesApp"
	"cards/restapidemo"
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
			4. Maps
			5. NotesApp - Interfaces demo
			6. API Route - Users using MUX
			7. GRPC demo
			8. GO Routine with channels
			9. GO Routine with wait Group
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
		case "4":
			mapsdemo.Execute()
		case "5":
			notesApp.Execute()
		case "6":
			restapidemo.InitialMigration()
			restapidemo.InitialRouter()
		case "7":
			grpcdemo.Execute()
		case "8":
			goroutinedemo.Execute()
		case "9":
			goroutinedemo.ExecuteWaitGroup()
		}

	}
}
