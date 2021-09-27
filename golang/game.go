package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const filePathWords string = "../words.txt"
const lives int = 10

func random_word_generator() string {
	var wordsString string = "apple pop snap cat rock"
	var words = strings.Fields(wordsString)
	var wordsListLength int = len(words)
	var randomNumber = rand.Intn(wordsListLength)

	// fmt.Printf("%v %T", words, words)

	var word string = words[randomNumber]
	return word
}

var word = random_word_generator()

func main() {
	fmt.Printf("%v", word)
}
