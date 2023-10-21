package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type MyData struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

var structData []MyData

func main() {
	// Open the text file for reading
	file, err := os.Open("../file2")
	if err != nil {
		fmt.Errorf("Failed to open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		data := MyData{
			Name: words[0],
			IP:   words[4],
		}
		structData = append(structData, data)
		//fmt.Printf("Struct data: %+v\n", structData

	}
	jsonData, err := json.Marshal(structData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//fmt.Println(string(jsonData))

	jsonFile, err := os.Create("../file4.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
