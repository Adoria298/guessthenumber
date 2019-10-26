package internal

// based on freshman.tech/golang-guess/

import (
	"fmt"
	"strconv"
)

// PlayerGuessMainLoop - main loop that allows the player to guess
func PlayerGuessMainLoop() {
	fmt.Println("Guess a number between 1 and 100.")

	secretNumber := GenerateRandomInteger(1, 100)

	var attempts int
	for {
		attempts++
		fmt.Println("Please input your guess.")

		input := ReadlnFromStdin("\r\n")

		guess, error := strconv.Atoi(input) // convert to integer
		if error != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			fmt.Println(error)
			continue
		}

		fmt.Println("Your guess is: ", guess)

		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number.")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number.")
		} else {
			fmt.Println("Correct. Game over.")
			break
		}
	}
}
