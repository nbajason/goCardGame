package main

import "fmt"

// type deck: slice of string
type deck []string

// d = first letter of deck
func (d deck) print() {
	fmt.Println("==========with index========")
	for i, card := range d {
		fmt.Println(i, card)
	}

}

func newDeck() deck {
	cards := deck{}
	suits := []string{"spades", "hearts", "diamonds", "clubs"}
	values := []string{"ace", "two", "three", "four"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}
