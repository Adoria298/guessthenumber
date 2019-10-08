package sharedfuncs

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

// GenerateRandomInteger - generates a random integer between min and max (inclusive)
func GenerateRandomInteger(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// ReadlnFromStdin - reads a line from os.Stdin and trims the delimeter passed.
func ReadlnFromStdin(delimeter string) string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSuffix(input, delimeter)
}
