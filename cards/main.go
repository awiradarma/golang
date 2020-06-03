package main

//import "fmt"

func main() {
	// var card string = "Ace of Spades"
	//cards := newDeck()
	
	//cards.saveToFile("my_cards")
	
	newcards := newDeckFromFile("my_cards")
	newcards.shuffle()
	newcards.print()
	//fmt.Println(newcards.toString())
	//hand, remainingCards:= deal(cards,5)
	//hand.print()
	//printStuff()
	//remainingCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

