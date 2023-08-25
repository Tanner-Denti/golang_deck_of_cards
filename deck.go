package main

import "fmt"

// Create a new type of 'deck', which is a slice of strings
type deck []string

// Generate and return a new 'deck'
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print our every card in the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i + 1, card)
	}
}
