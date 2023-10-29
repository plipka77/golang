package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var CSVFILE string = "./resource/cars.csv"

func main() {
	_, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println("Can't find file")
	} else {
		fmt.Println("Found the file")
	}

	file, err := os.Open(CSVFILE)
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error reading csv file")
		return
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
