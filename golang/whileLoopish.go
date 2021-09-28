// Golang program to show how
// to take input from the user
package main

import "fmt"

func game() string {

	var first string

	// Taking input from user
	fmt.Printf("press q to quit\n")
	fmt.Scanln(&first)

	return first
}

// main function
func main() {

	// var play string = game()

	for {
		var play string = game()
		if play == string('q') {
			break
		}
		if play != string('q') {
			fmt.Printf("lets play :)\n")
		}
	}

}
