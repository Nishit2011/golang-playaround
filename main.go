package main

func main() {
	//:= is used only when we declare the variable
	// if

	//cards := newDeck()
	//cards = append(cards, "Six of Spades")

	//since cards is declared above to be
	//of deck type,,it gets access to the print
	//method declared inside deck
	//cards.print()

	//since both main.go and deck.go are inside the same
	//package we dont need to import any method from deck.go
	//we can directly invoke the deal method here
	//for ex:
	//deal(cards, 5)

	//since deal returns two values
	//first val will be assigned to hand and the
	//second to remainingCards
	// hand, remainingCards := deal(cards, 5)

	// //printing both the values
	// hand.print()
	// remainingCards.print()

	//changing the slice into string
	//fmt.Println(cards.toString())

	//saving list of cards slice in a file named my_cards
	//cards.saveToFile("my_cards")

	//reading the string from file my_cards and converting them back to splice of type deck
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
