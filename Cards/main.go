package main

func main() {
	cards := newDeckFromFile("deck.txt")
	cards.print()

	cards.shuffle()
	cards.print()
}

func assignNewCard(newCardText string) string {
	return newCardText
}
