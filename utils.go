package main

import "fmt"


func FormatPrint(data WordInfo){

	// ensure word info is not empty
	if data.Word == "" {
		fmt.Println("No data found")
		return
	}

	fmt.Println("Word:", data.Word)
	fmt.Println("Phonetic:", data.Phonetic)
	fmt.Println("Phonetics:")
	for _, phonetic := range data.Phonetics {
		fmt.Printf("  Text: %s, Audio: %s\n", phonetic.Text, phonetic.Audio)
	}
	fmt.Println("Meanings:")
	for _, meaning := range data.Meanings {
		fmt.Printf("  Part of Speech: %s\n", meaning.PartOfSpeech)
		for _, definition := range meaning.Definitions {
			fmt.Printf("    Definition: %s\n", definition.Definition)
		}
	}
	fmt.Println("Source URLs:", data.SourceUrls)
}

func contains(slice []string, item string) bool {
    for _, element := range slice {
        if element == item {
            return true
        }
    }
    return false
}