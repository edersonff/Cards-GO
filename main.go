package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	cards.saveToFile("cards.txt")

	cards = newDeckFromFile("cards.txt")
	cards.print()
}
