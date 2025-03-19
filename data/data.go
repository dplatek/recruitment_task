package data

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadDataFromFile(filename string) ([]int, error) {
	log.Println("Info: Loading data from file:", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error: Error loading file %s: %v", filename, err)
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	var numbers []int

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("Error: Invalid number in file: %s", line)
			return nil, err
		}
		numbers = append(numbers, number)
	}

	log.Printf("Info: Successfully loaded %d records from file", len(numbers))
	return numbers, nil
}
