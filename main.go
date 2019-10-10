package main

import (
	"fmt"

	"github.com/adoria298/guessthenumber/sharedfuncs"

	"github.com/adoria298/guessthenumber/guessing"
)

func main() {
	fmt.Println("Would you like to guess the number, or shall I?")
	fmt.Println("Input 'Y' if you would, and 'N' if I should.")

	input := sharedfuncs.ReadlnFromStdin("\r\n")

	if input == "Y" {
		guessing.PlayerGuessMainLoop()
	} else if input == "N" {
		guessing.AIGuessMainLoop()
	} else {
		fmt.Println("I don't understand this input. Exiting.")
	}
}
