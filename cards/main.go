package main

// import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

}
