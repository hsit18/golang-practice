package main

import (
	"fmt"
	"main/deck"
	"time"
)

func main() {
	fmt.Println("Starting main....", time.Now())
	cards := deck.NewDeck()
	hand, remainingDeck := deck.Deal(cards, 5)
	hand.Print()
	fmt.Println("remainingDeck....")
	remainingDeck.Print()
	remainingDeck.SaveToFile("cards.txt")
}

func newCard() string {
	return "Five of Diamonds"
}
