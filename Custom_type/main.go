package main

import "fmt"

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// fmt.Println("hand of cards are:")
	// hand.print()
	// fmt.Println("remaining cards are:")
	// remainingCards.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("myCards")
	// newDeckFromFile("myCards")

	cards := newDeck()
	cards.print()
	fmt.Println("")
	fmt.Println("")
	cards.shuffle()
	cards.print()

}
