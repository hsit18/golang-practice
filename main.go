package main

import "fmt"

func main() {
	fmt.Println("Starting main....");
	cards := deck{"Ace Cards", newCard()}
	cards = append(cards, "Six of Shades")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds";
}