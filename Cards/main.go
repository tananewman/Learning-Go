package main

func main() {
	cards := getNewDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}

func assignNewCard(newCardText string) string {
	return newCardText
}
