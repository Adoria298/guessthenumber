package internal

import (
	"fmt"
)

// AIGuessMainLoop - main loop for AI version of the guessing game.
func AIGuessMainLoop() {
	fmt.Println("Think of a number between 1 and 100.")
	min := 1
	max := 100

	var attempts int
	for {
		attempts++
		guess := GenerateRandomInteger(min, max)
		fmt.Println("Is this your number:", guess, "?")
		fmt.Println("Input 'h' if it is higher than your number,",
			"'e' if it is equal to it, and 'l' if it is lower.")
		guessCorrect := ReadlnFromStdin("\r\n")
		if guessCorrect == "e" {
			fmt.Println("Yay! I got it!")
			break
		} else if guessCorrect == "l" {
			min = guess + 1
		} else if guessCorrect == "h" {
			max = guess - 1
		} else {
			fmt.Println("Oh well! That's a shame!")
		}
	}
}
