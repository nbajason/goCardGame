package main

import "fmt"

// var card string = "test"
// array: fixed
// sclice: grow / shrink
var cards deck

func main() {
	// *** manually create new card deck
	// cardList = append(cardList, "Ace of Diamonds")
	// cardList = append(cardList, "Six Diamonds")
	// cardList = append(cardList, newCard())
	// basic type: bool string int float64

	// *** deal
	cards = newDeck()
	handCards, remainingCards := deal(cards, 5)

	handCards.print()
	remainingCards.print()

	// ** convert deck to string
	fmt.Println(handCards.toString())
	// fmt.Println([]byte(cards.toString()))

	// ** save to file
	cards.saveToFile("myCards.txt")
	cards.saveToFile("remainingCards.txt")
}
