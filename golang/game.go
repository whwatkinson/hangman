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

func main() {
	var word = random_word_generator()
	var wordLength int = len(word)
	var word_chars = make([]string, wordLength)

	fmt.Printf("%v\n", word_chars)

	for i := 0; i < wordLength; i++ {
		var underScore string = string('_')
		word_chars[i] = underScore
	}

	fmt.Printf("%v\n", word_chars)

	var allLetters string = "abcdefghijklmnopqrstuvwxyz-"
	var lettersRemaining = make([]string, len(allLetters))

	for i := 0; i < len(allLetters); i++ {
		lettersRemaining[i] = string(allLetters[i])
	}

	fmt.Printf("%v", lettersRemaining)
}
