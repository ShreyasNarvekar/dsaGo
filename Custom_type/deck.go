package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string //creating custom type deck of stringSlice.

func (d deck) print() { //Printing the cards in the deck..
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck { //Creating a new deck of cards.
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suits)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) { //dealing the card
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string { //type casting the deck type to stringSlice nd converting into one string.
	//passing the deck converted to type slice of
	// string and passing  ","  for
	// the seperator(which seprate everything with comma).
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error { //Saving deck to the hardisk
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck { // NEW Deck of Cards From the File method
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source:= rand.NewSource(time.Now().UnixNano())
	r:= rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
