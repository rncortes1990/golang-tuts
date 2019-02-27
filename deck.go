package main

import (
	"fmt"
	"strings"
)

type deck [] string

func newDeck() deck{

	cards := deck {}
	cardSuits := [] string {"Spades","Diamonds","Hearts","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	for _,suit := range cardSuits{
		for _,value := range cardValues{
			cards = append(cards, value  +" of "+ suit )
		}
	}
	return cards
}

func (d deck) print(){
	for i,card := range d {
		fmt.Println(i,card)
	}
}

func  deal(d deck,handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}
func (d deck) toString() string{
	//type conversion
	return strings.Join([] string (d),",")

}
func (s string) toByteSlice () [] byte{
	
}
/* FUNCTION DEFINITION 
*
* function (arguments) func_name() return_value{
*   code goes here
*  }	
*/