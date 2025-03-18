package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Function to load data from a txt file line by line and convert each line to an integer
func loadDataFromFile(filename string) ([]int, error) {
	var numbers []int

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert each line to an integer
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			// If conversion fails, skip this line
			fmt.Println("Skipping invalid line:", line)
			continue
		}
		numbers = append(numbers, num)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

// Function to initialize the server and load the data
func initializeServer() ([]int, error) {
	// Load data from a file before starting the server
	numbers, err := loadDataFromFile("input.txt")
	if err != nil {
		return nil, err
	}

	return numbers, nil
}

func endpointHandler(c *gin.Context) {
	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid index, must be an integer")
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("Success: You reached /endpoint/%d", index))
}

func main() {
	// Initialize the server and load data from the file
	numbers, err := initializeServer()
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}

	// Print the numbers (optional)
	fmt.Println("Loaded numbers from file:", numbers)

	// Initialize the Gin router
	r := gin.Default()
	r.GET("/endpoint/:index", endpointHandler)

	// Start the Gin server
	fmt.Println("Server is running on :8080")
	r.Run(":8080")
}
