package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const filePathWords string = "../words.txt"

var lives int = 10
var playing bool = true

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

func stringInSlice(
	a string,
	slice []string,
) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}

func stringInSliceIndexes(
	playerGuess string,
	wordChars []string,
) []int {
	var indexes []int
	for i, b := range wordChars {
		if b == playerGuess {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func removeStringInSlice(
	playerGuess string,
	slice []string,
) []string {
	for i, b := range slice {
		if b == playerGuess {
			slice[i] = " "
		}
	}
	return slice
}

func preCheckScreen(
	playerGuess string,
	rightGuess []string,
	wrongGuess []string,
	lettersRemaining []string,
) bool {

	// Same letter already guess and right
	if stringInSlice(playerGuess, rightGuess) {
		fmt.Printf("Already guessed, %v it was there\n", playerGuess)
		return false
	}

	// Same letter already guess and wrong
	if stringInSlice(playerGuess, wrongGuess) {
		fmt.Printf("Already guessed, %v it was not there\n", playerGuess)
		return false
	}

	// Not a recgnised charater
	if stringInSlice(playerGuess, lettersRemaining) != true {
		fmt.Printf("input not recognised.\n")
		return false
	}

	return true

}

func wrongGuess(
	playerGuess string,
	lettersRemaining []string,
	wrongGuesses []string,
) int {
	fmt.Printf("%v was not in the word\n", playerGuess)

	for i, b := range lettersRemaining {
		if b == playerGuess {
			lettersRemaining[i] = " "
		}
	}
	lettersRemaining = removeStringInSlice(playerGuess, lettersRemaining)

	wrongGuesses = append(wrongGuesses, playerGuess)
	return -1
}

func rightGuess(
	playerGuess string,
	wordChars []string,
	lettersRemaining []string,
	wordBlank []string,
	rightGuesses []string,
) int {

	fmt.Printf("%v was in the word\n", playerGuess)
	lettersRemaining = removeStringInSlice(playerGuess, lettersRemaining)

	var indexes []int = stringInSliceIndexes(playerGuess, wordChars)

	for _, pos := range indexes {
		wordBlank[pos] = playerGuess
	}

	rightGuesses = append(rightGuesses, playerGuess)

	return 0
}

func guessCheck(
	playerGuess string,
	wordChars []string,
	lettersRemaining []string,
	wordBlank []string,
	rightGuesses []string,
	wrongGuesses []string,
) int {
	var livesToRemove int = 0
	if stringInSlice(playerGuess, wordChars) {
		livesToRemove = rightGuess(
			playerGuess,
			wordChars,
			lettersRemaining,
			wordBlank,
			rightGuesses,
		)
	} else {
		livesToRemove = wrongGuess(
			playerGuess,
			lettersRemaining,
			wrongGuesses,
		)
	}
	return livesToRemove
}

func main() {
	// get word
	var word = random_word_generator()
	var wordLength int = len(word)
	var wordChars = make([]string, wordLength)
	for i := 0; i < wordLength; i++ {
		wordChars[i] = string(word[i])
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
	var roundNumber int = 0
	var rightGuesses = make([]string, 27)
	// var wrongGuesses = make([]string, 27)
	var wrongGuesses []string

	fmt.Printf("******************\n")
	fmt.Printf("Welcome to Hangman\n")
	fmt.Printf("******************\n")

	// This is now the while loop
	for lives > 0 {

		// VICTORY CHECK
		if stringInSlice(string("_"), wordBlank) != true {
			fmt.Printf("Congratulations, you won!\n")
			break
		}
		// ROUND SETUP
		roundNumber += 1
		fmt.Printf("\n\n\n\n")
		fmt.Printf("********\n")
		fmt.Printf("Round %v\n", roundNumber)
		fmt.Printf("********\n\n")

		// TODO board of the gallows
		fmt.Printf("You have %v live(s) remaining\n\n", lives)
		fmt.Printf("word:      %v\n\n", wordBlank)
		fmt.Printf("wrong:     %v\n\n", wrongGuesses)
		fmt.Printf("letters:   %v\n\n", lettersRemaining)

		// PLAYER INPUT
		var playerGuess string
		fmt.Printf("have a guess --->: ")

		fmt.Scanln(&playerGuess)
		fmt.Printf("\n")

		var ok bool = preCheckScreen(playerGuess, rightGuesses, wrongGuesses, lettersRemaining)

		if ok {
			lives += guessCheck(
				playerGuess,
				wordChars,
				lettersRemaining,
				wordBlank,
				rightGuesses,
				wrongGuesses,
			)
		}

		time.Sleep(2 * time.Second)
	}

	fmt.Printf("The word I was thinking of was: %v\n", word)
}
