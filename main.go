package main

import "fmt"

func main() {

	cards := newDeck()
	cards.shuffleDeck()
	hand, _ := deal(cards, 5)
	cards.saveToFile("myCards")
	for _, s := range hand {
		fmt.Println(s)
	}
}
