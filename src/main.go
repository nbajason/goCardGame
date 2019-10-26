package main

import "fmt"

// var card string = "test"
// array: fixed
// sclice: grow / shrink
var handCards deck
var remainingCards deck

func main() {
	// *** manually create new card deck
	// cardList = append(cardList, "Ace of Diamonds")
	// cardList = append(cardList, "Six Diamonds")
	// cardList = append(cardList, newCard())
	// basic type: bool string int float64

	handCards = newDeckFromFile("myCards.txt")
	remainingCards = newDeckFromFile("remainingCards.txt")

	// *** deal
	handCards, remainingCards := deal(handCards, 5)

	// *** read cards from file

	handCards.print()
	remainingCards.print()

	// ** convert deck to string
	fmt.Println(handCards.toString())
	// fmt.Println([]byte(cards.toString()))

	// ** save to file
	handCards.saveToFile("myCards.txt")
	handCards.saveToFile("remainingCards.txt")
}
