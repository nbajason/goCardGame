package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
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

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func newDeckFromFile(fileName string) deck {
	var result deck
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		// return new deck or exit
		// os.Exit(1)
		result = newDeck()
	} else {
		// convert byte slice to string
		str := string(bs)
		// convert string(seperate by comma) to deck
		strSlice := strings.Split(str, ",")
		result = deck(strSlice)
	}
	return result
}
