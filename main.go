package main

import (
	"fmt"
	"hangman/hangman"
)

func main() {
	g := hangman.New(8, "golang")
	fmt.Println(g)
}
