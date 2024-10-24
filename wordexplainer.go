package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func explain(word string) WordInfo {

	fmt.Println("-------🔖-----------------------------------✍️--------")


	fmt.Printf("This is what i found for the word %s 😇 \n", word)

	fmt.Println("-------🔖-----------------------------------✍️--------")

	word = strings.TrimSpace(word)
	baseURL := "https://api.dictionaryapi.dev/api/v2/entries/en/"
	url := baseURL + word
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Error fetching data:", err)
		return WordInfo{}
	}

	defer resp.Body.Close()

	var data []WordInfo // The response is an array of WordInfo objects
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error reading data:", err)
		return WordInfo{}
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal("Error parsing data:", err)
		return WordInfo{}
	}

	// Since it's an array, return the first entry
	return data[0]
}