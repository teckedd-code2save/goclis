package main

// WordInfo is a struct that holds the information of a word
type Phonetic struct {
	Text     string `json:"text"`
	Audio    string `json:"audio"`
}

type Definition struct {
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
}

type Meaning struct {
	PartOfSpeech string      `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type WordInfo struct {
	Word      string     `json:"word"`
	Phonetic  string     `json:"phonetic"`
	Phonetics []Phonetic `json:"phonetics"`
	Meanings  []Meaning  `json:"meanings"`
	SourceUrls []string  `json:"sourceUrls"`
}
