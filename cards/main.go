package main

func main() {
	cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// fmt.Println(hand.toString())
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
