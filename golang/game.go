package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const filePathWords string = "../words.txt"

func random_word_generator() string {
	var wordsString string = "apple pop snap cat rocks tiple"
	var words = strings.Fields(wordsString)
	var wordsListLength int = len(words)

	rand.Seed(time.Now().UnixNano())
	var randomNumber = rand.Intn(wordsListLength)
	var word string = words[randomNumber]

	// fmt.Printf("%v %T\n", randomNumber, randomNumber)
	return word
}

func main() bool {
	// get word
	var word = random_word_generator()
	var wordLength int = len(word)
	var word_chars = make([]string, wordLength)
	for i := 0; i < wordLength; i++ {
		word_chars[i] = string(word[i])
	}

	// create word blank
	var wordBlank = make([]string, wordLength)
	for i := 0; i < wordLength; i++ {
		wordBlank[i] = string('_')
	}

	// get remaining letters
	var allLetters string = "abcdefghijklmnopqrstuvwxyz-"
	var lettersRemaining = make([]string, len(allLetters))
	for i := 0; i < len(allLetters); i++ {
		lettersRemaining[i] = string(allLetters[i])
	}

	// set up game and counters
	// var roundNumber int = 0
	var lives int = 10
	// var rightguess = make([]string, 27)
	// var wrongGuess = make([]string, 27)

	fmt.Printf("******************")
	fmt.Printf("Welcome to Hangman")
	fmt.Printf("******************\n")

	if lives < 0 {
		fmt.Printf("The word I was thinking of was: %v", word)
	} else {
		fmt.Printf("Lives remaing: %v", lives)
		fmt.Printf("*********")
		fmt.Printf("GAME OVER")
		fmt.Printf("*********\n")
		return false
	}

	return true
}
