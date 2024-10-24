package main

import (
	"fmt"
)

func interact() {
	var name string
	var hobby string

	fmt.Print("Whats your name: ")
	fmt.Scanln(&name)

	fmt.Print("Whats your hobby: ")
	fmt.Scanln(&hobby)
	fmt.Println("************************************************************")
	fmt.Printf("* Hello %s,nice to meet you. Guess I also love %s \n", name, hobby)
	fmt.Println("*************************************************************")
	
}