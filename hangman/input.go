package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false
	for valid != true {
		fmt.Print("Choose a letter")
		guess, err := reader.ReadString('\n')
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess)
		if len(guess) != 1 {
			fmt.Print("You just need to insert one char")
			continue
		}
		valid = true
	}
	return guess, nil
}
