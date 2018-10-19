//in this file we are creating a new type
//with a function that has a receiver

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

//creating an array of strings and add a bunch
//of functions specifically made to work with it

//create a new type of 'deck'
//which is a slice of strings

type deck []string

//in golang the return type of the
//function has to be added along with
//function declaration
//below function will return values with deck type
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// _ represents values that arent used
	//was used instead of convemtional i and j as i and j are never used

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+"of"+suit)
		}
	}

	return cards
}

//any variable of type "deck" gets access to the
//print method

//d is the actual copy of the deck we're working
//with

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//handSize is the argument being passed to the function which is of int data type
//two values of deck type will be returned
func deal(d deck, handSize int) (deck, deck) {
	//first value returned will start from beginning of slice to the handSize passed
	//second value returned will start from handsize value to the end

	return d[:handSize], d[handSize:]
}

//converting a deck into a slice of string
func (d deck) toString() string {
	//strings.join takes slice of string and converts it into string separated by comma

	return strings.Join([]string(d), ",")
}

//function to save a file to hard drive
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//reads the filename named by filename and returns the content, here deck
func newDeckFromFile(filename string) deck {
	//bs is byte slice
	//err is value of type 'error
	//if nothing went wrong, it will have a value of 'nil'
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//Option #1 - log the error and return a call to newDeck()
		//Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		//exit the program
		os.Exit(1)
	}

	//in split we pass string and separate it with comma
	//convert into a splice, which in next line we use to turn into type deck
	s := strings.Split(string(bs), ",")
	return deck(s)
}

//func to shuffle a deck of cards
//has a receiver(d deck) so that we can do cards.shuffle
func (d deck) shuffle() {
	for i := range d {
		// generates a random number using library rand method
		//random number is gen b/w length of deck of cards-1 and 0
		newPosition := rand.Intn(len(d) - 1)

		//swapping the card at randomPosition and current position
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}