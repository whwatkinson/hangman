package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

const filePathWords string = "words.txt"

func main() {
	data, err := ioutil.ReadFile(filePathWords)
	if err != nil {
		fmt.Println("File reading error", err)
	}

	var words = strings.Fields(string(data))
	var wordsListLength int = len(words)

	rand.Seed(time.Now().UnixNano())
	var randomNumber = rand.Intn(wordsListLength)
	var word string = words[randomNumber]

	fmt.Printf("%v", word)
}
