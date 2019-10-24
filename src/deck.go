package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// type deck: slice of string
type deck []string

// d = first letter of deck
func (d deck) print() {
	fmt.Println("==========deck with index========")
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	deckString := strings.Join([]string(d), ",")
	return deckString
	// sDeck := fmt.SPrintf(deck)
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
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

func (d deck) saveToFile(fileName string) error {
	str := d.toString()
	return ioutil.WriteFile(fileName, []byte(str), 0666)
}

func shuffle() {

}
