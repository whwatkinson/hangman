package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

const filePathWords string = "/words.txt"

var lives int = 10

func randomWordGenerator(filePath string) string {

	pwd, _ := os.Getwd()
	data, err := ioutil.ReadFile(pwd + filePath)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	rand.Seed(time.Now().UnixNano())
	var words = strings.Fields(string(data))
	var word string = words[rand.Intn(len(words))]

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
	lettersRemaining []string,
) []string {
	for i, b := range lettersRemaining {
		if b == playerGuess {
			lettersRemaining[i] = " "
		}
	}
	return lettersRemaining
}

func preCheckScreen(
	playerGuess string,
	rightGuesses []string,
	wrongGuesses []string,
	lettersRemaining []string,
) bool {

	// Same letter already guess and right
	if stringInSlice(playerGuess, rightGuesses) {
		fmt.Printf("Already guessed %v, it was there\n", playerGuess)
		return false
	}

	// Same letter already guess and wrong
	if stringInSlice(playerGuess, wrongGuesses) {
		fmt.Printf("Already guessed %v, it was not there\n", playerGuess)
		return false
	}

	// Not a recgnised charater
	if !stringInSlice(playerGuess, lettersRemaining) {
		fmt.Printf("input not recognised.\n")
		return false
	}

	return true

}

func wrongGuess(
	playerGuess string,
	lettersRemaining []string,
	wrongGuesses []string,
) (int, []string) {
	fmt.Printf("%v was not in the word\n", playerGuess)

	for i, b := range lettersRemaining {
		if b == playerGuess {
			lettersRemaining[i] = " "
		}
	}
	lettersRemaining = removeStringInSlice(playerGuess, lettersRemaining)

	wrongGuesses = append(wrongGuesses, playerGuess)
	return -1, wrongGuesses
}

func rightGuess(
	playerGuess string,
	wordChars []string,
	lettersRemaining []string,
	wordBlank []string,
	rightGuesses []string,
) (int, []string) {

	fmt.Printf("%v was in the word\n", playerGuess)
	lettersRemaining = removeStringInSlice(playerGuess, lettersRemaining)
	var indexes []int = stringInSliceIndexes(playerGuess, wordChars)

	for _, pos := range indexes {
		wordBlank[pos] = playerGuess
	}

	rightGuesses = append(rightGuesses, playerGuess)

	return 0, rightGuesses
}

func guessCheck(
	playerGuess string,
	wordChars []string,
	lettersRemaining []string,
	wordBlank []string,
	rightGuesses []string,
	wrongGuesses []string,
) (int, []string, []string) {
	var livesToRemove int = 0
	if stringInSlice(playerGuess, wordChars) {
		livesToRemove, rightGuesses = rightGuess(
			playerGuess,
			wordChars,
			lettersRemaining,
			wordBlank,
			rightGuesses,
		)
	} else {
		livesToRemove, wrongGuesses = wrongGuess(
			playerGuess,
			lettersRemaining,
			wrongGuesses,
		)
	}
	return livesToRemove, rightGuesses, wrongGuesses
}

func main() {
	// get word
	var word = randomWordGenerator(filePathWords)
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
	var allLetters string = "abcdefghijklmnopqrstuvwxyz"
	var lettersRemaining = make([]string, len(allLetters))
	for i := 0; i < len(allLetters); i++ {
		lettersRemaining[i] = string(allLetters[i])
	}

	// set up game and counters
	var roundNumber int = 0
	var rightGuesses []string
	var wrongGuesses []string
	// var wrongGuesses = make([]string, 27)

	fmt.Printf("******************\n")
	fmt.Printf("Welcome to Hangman\n")
	fmt.Printf("******************\n")

	// This is now the while loop
	for {

		// VICTORY CHECK
		if !stringInSlice(string("_"), wordBlank) {
			fmt.Printf("\n\n\n*************************\n")
			fmt.Printf("Congratulations, you won!\n")
			fmt.Printf("The word I was think of was %v\n", word)
			fmt.Printf("*************************\n")
			break
		}

		// GAME OVER CHECK
		if lives == 0 {
			fmt.Printf("\n\n\n*********\n")
			fmt.Printf("GAME OVER\n")
			fmt.Printf("*********\n")
			fmt.Printf("The word I was thinking of was: %v\n", word)

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
		fmt.Printf("word:                %v\n\n", wordBlank)
		fmt.Printf("wrong guesses:       %v\n\n", wrongGuesses)
		fmt.Printf("letters remaining:   %v\n\n", lettersRemaining)

		// PLAYER INPUT
		var playerGuess string
		fmt.Printf("have a guess --->: ")

		fmt.Scanln(&playerGuess)
		fmt.Printf("\n")

		var ok bool = preCheckScreen(playerGuess, rightGuesses, wrongGuesses, lettersRemaining)

		if ok {
			var livesToRemove int

			livesToRemove, rightGuesses, wrongGuesses = guessCheck(
				playerGuess,
				wordChars,
				lettersRemaining,
				wordBlank,
				rightGuesses,
				wrongGuesses,
			)
			lives += livesToRemove
		}

		time.Sleep(1 * time.Second)
	}
}
