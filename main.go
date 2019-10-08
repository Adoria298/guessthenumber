package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Guess a number between 1 and 100.")
	fmt.Println("Please input your guess.")

	secretNumber := generateRandomInteger(1, 100)
	fmt.Println("The secret number is: ", secretNumber)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\r\n")

	guess, error := strconv.Atoi(input) // convert to integer
	if error != nil {
		fmt.Println("Invalid input. Please enter an integer.")
		fmt.Println(error)
		return
	}

	fmt.Println("Your guess is: ", guess)

	if guess > secretNumber {
		fmt.Println("Your guess is bigger than the secret number.")
	} else if guess < secretNumber {
		fmt.Println("Your guess is smalled than the secret number.")
	} else {
		fmt.Println("Correct. Game over.")
	}
}

func generateRandomInteger(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
