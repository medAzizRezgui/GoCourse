package main

import "fmt"

func main() {
	//cards := newDeck()
	////hand, remainingDeck := deal(cards, 5)
	//cards.saveToFile("myCards")

	cards := readDeckFromFile("myCards")
	for _, card := range cards {
		fmt.Println(card)
	}
}
