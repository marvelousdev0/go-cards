package main

func main() {
	/* cards := newDeck()
	hand, remaining := cards.deal(5)
	hand.print()
	remaining.print()

	cardString := cards.toString()
	fmt.Println(cardString)

	cards.saveToFile("my_cards") */
	cards := newDeckFromFile("my_cards")
	cards.print()
	cards.shuffle()
	cards.print()
}
