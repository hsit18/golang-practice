package main

import (
	"cards/deck"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting main....", time.Now())
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
