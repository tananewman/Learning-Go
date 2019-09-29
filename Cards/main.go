package main

func main() {
	cards := getNewDeckFromFile("deck.txt")
	cards.print()

	cards.shuffle()
	cards.print()
}

func assignNewCard(newCardText string) string {
	return newCardText
}
