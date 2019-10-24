package main

import "fmt"

// var card string = "test"
// array: fixed
// sclice: grow / shrink
var cards deck

func main() {
	// *** new card deck
	// cardList = append(cardList, "Ace of Diamonds")
	// cardList = append(cardList, "Six Diamonds")
	// cardList = append(cardList, newCard())
	// basic type: bool string int float64
	// var car string = ""

	// *** deal
	// cards = newDeck()
	// handCards, remainingCards := deal(cards, 5)

	// handCards.print()
	// remainingCards.print()

	// ** convert deck to string
	cards := newDeck()
	fmt.Println(cards.toString())
	// fmt.Println([]byte(cards.toString()))
	cards.saveToFile("myCards.txt")
}
