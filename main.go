package main

import "fmt"

func main(){
	

	// cards := deck {"ace of spades", newCard()}
	cards := newDeck()
	// cards = append (cards,"six of spades")
	hand,deckRest:= deal(cards,5)
	fmt.Println("--mazo--")
	cards.print()
	fmt.Println("--mano jugador--")
	hand.print()
	fmt.Println("--resto de mazo--")
	deckRest.print()

}

func newCard () string{
	return "five of diamondsdiamonds"
}
