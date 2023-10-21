package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type data struct {
	Count   int     `json:"count"`
	Entries Entries `json:"entries"`
}

type Entries []struct {
	API         string `json:"API"`
	Description string `json:"Description"`
	Auth        string `json:"Auth"`
	HTTPS       bool   `json:"HTTPS"`
	Cors        string `json:"Cors"`
	Link        string `json:"Link"`
	Category    string `json:"Category"`
}

func main() {

	// Take input from the user.
	var keyword string
	fmt.Println("Enter the input:")
	_, err := fmt.Scanln(&keyword)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Open the file
	file, err := os.Open("../problem2.json")
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}
	defer file.Close()

	// Decode the json file
	decoder := json.NewDecoder(file)

	var structdata data

	// Convert json data to struct
	err = decoder.Decode(&structdata)
	if err != nil {
		fmt.Println("Error decoding data", err)
		return
	}

	// Check if the struct has data that matches user input
	for _, entry := range structdata.Entries {
		if entry.API == keyword || entry.Description == keyword {
			fmt.Println(entry.API)
		}
	}
}
