package guessing

// based on freshman.tech/golang-guess/

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// PlayerGuessMainLoop - main loop that allows the player to guess
func PlayerGuessMainLoop() {
	fmt.Println("Guess a number between 1 and 100.")

	secretNumber := generateRandomInteger(1, 100)

	var attempts int
	for {
		attempts++
		fmt.Println("Please input your guess.")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\r\n")

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

func generateRandomInteger(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
