package main

import "fmt"

func main() {
	//var cards string = "Ace of spaces"
	cards := []string{newCard(), "heart", "diamond"}
	cards = append(cards, "Ace of diamonds")
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)

}

func newCard() string {
	return "Five of diamonds"
}
