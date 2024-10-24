package main

import (
	"fmt"
	"bufio"
	"os"

)

func guessworrd() string{
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Guess a word: ")

	word , _ := reader.ReadString('\n')

	fmt.Printf("You guessed the word %s\n", word)	
    
	return word

}