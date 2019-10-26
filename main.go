package main

import (
	"fmt"

	"github.com/adoria298/guessthenumber/internal"
)

func main() {
	fmt.Println("Would you like to guess the number, or shall I?")
	fmt.Println("Input 'Y' if you would, and 'N' if I should.")

	input := internal.ReadlnFromStdin("\r\n")

	if input == "Y" {
		internal.PlayerGuessMainLoop()
	} else if input == "N" {
		internal.AIGuessMainLoop()
	} else {
		fmt.Println("I don't understand this input. Exiting.")
	}
}
