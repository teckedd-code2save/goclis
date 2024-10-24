package main

import (
	"os"
	"fmt"
)


func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("-------------------------------------------------------------------------")

		fmt.Println("| Hi buddy, I am a CLI app. I can chat with you or you can guess a word |")

		fmt.Println("-------------------------------------------------------------------------")
		fmt.Println("Just enter 'chat' or 'guess' for some thrill")
		
	} else {
		fmt.Print("You want to ")
		for i,arg := range args {

			if i !=0 {
				fmt.Print("and ")
			}
			fmt.Printf("%s\n", arg)
		}

		if contains(args, "chat" ){
			interact()
		} 
		if contains(args, "guess") {
			word := guessworrd()
			data := explain(word)
			formatPrint(data)
		}

	}
}