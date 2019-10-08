package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guess a number between 1 and 100.")
	fmt.Println("Please input your guess.")

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
}
