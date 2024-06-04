package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/hsit18/golang-practice/channeldemo"
	"github.com/hsit18/golang-practice/contextdemo"
	"github.com/hsit18/golang-practice/deck"
	"github.com/hsit18/golang-practice/goroutinedemo"
	"github.com/hsit18/golang-practice/grpcdemo"
	"github.com/hsit18/golang-practice/mapsdemo"
	"github.com/hsit18/golang-practice/notesApp"
	"github.com/hsit18/golang-practice/restapidemo"
	"github.com/hsit18/golang-practice/structsdemo"
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
			10. Context demo
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
		case "10":
			contextdemo.Execute()
		}

	}
}
