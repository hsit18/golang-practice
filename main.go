package main

import "fmt"

func main() {
	fmt.Println("Starting main....")
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	fmt.Println("remainingDeck....")
	remainingDeck.print()
	remainingDeck.saveToFile("cards.txt")
}

func newCard() string {
	return "Five of Diamonds"
}
