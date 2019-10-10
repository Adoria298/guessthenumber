package guessing

import (
	"fmt"

	"github.com/adoria298/guessthenumber/sharedfuncs"
)

// AIGuessMainLoop - main loop for AI version of the guessing game.
func AIGuessMainLoop() {
	fmt.Println("Think of a number between 1 and 100.")

	var attempts int
	for {
		attempts++
		guess := sharedfuncs.GenerateRandomInteger(1, 100)
		fmt.Println("Is this your number:", guess, "?")
		fmt.Println("Input 'h' if it is higher than your number,",
			"'e' if it is equal to it, and 'l' if it is lower.")
		guessCorrect := sharedfuncs.ReadlnFromStdin("\r\n")
		if guessCorrect == "e" {
			fmt.Println("Yay! I got it!")
			break
		} else {
			fmt.Println("Oh well! That's a shame!")
		}
	}
}
