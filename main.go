package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards.csv")

	cards_test := newDeckFromFile("my_cards.csv")
	cards_test.print()
}
