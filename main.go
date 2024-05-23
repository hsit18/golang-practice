package main

import (
	"bufio"
	"cards/channeldemo"
	"cards/deck"
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
	`)

	for scanner.Scan() {
		fmt.Println("You selected: ", scanner.Text())
		fmt.Println("*************************************************************************************************************")

		switch scanner.Text() {
		case "1":
			channeldemo.ExecuteFibonacci()
		case "2":
			cards := deck.NewDeck()
			cards.Shuffle()
			fmt.Println("cards suffled....")
			cards.Print()

			hand, remainingDeck := deck.Deal(cards, 5)
			hand.Print()
			//fmt.Println("remainingDeck....")
			//remainingDeck.Print()
			remainingDeck.SaveToFile("cards.txt")

			deckContent := deck.NewDeckFromFile("cards.txt")
			fmt.Println("deckContent....")
			deckContent.Print()
		}

	}
}
