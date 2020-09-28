package main

import (
	"fmt"
	"hangman/hangman"
	"os"
)

func main() {
	g := hangman.New(8, "golang")
	fmt.Println(g)

	l, err := hangman.ReadGuess()
	if err != nil {
		fmt.Printf("Cannot read the input")
		os.Exit(1)
	}
	fmt.Print(l)
}
