package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Function to load data from a txt file into a slice of integers
func loadDataFromFile(filename string) ([]int, error) {
	// Use os.ReadFile to read the contents of the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Split the data by new lines and convert each line to an integer
	lines := strings.Split(string(data), "\n")
	var numbers []int

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid number in file: %s", line)
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

// Function to initialize the server and load the data
func initializeServer() ([]int, error) {
	// Load data from a file before starting the server
	fileData, err := loadDataFromFile("input.txt")
	if err != nil {
		return nil, err
	}

	return fileData, nil
}

// Endpoint handler to get the index of a value in the slice
func endpointHandler(c *gin.Context, data []int) {
	valueStr := c.Param("value")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid value, must be an integer")
		return
	}

	// Search for the value in the slice
	for i, v := range data {
		if v == value {
			c.String(http.StatusOK, fmt.Sprintf("Value %d found at index %d", value, i))
			return
		}
	}

	// If the value is not found
	c.String(http.StatusNotFound, fmt.Sprintf("Value %d not found", value))
}

func main() {
	// Initialize the server and load data from the file
	data, err := initializeServer()
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}

	// Initialize the Gin router
	r := gin.Default()

	// Define the route with the handler
	r.GET("/endpoint/:value", func(c *gin.Context) {
		endpointHandler(c, data)
	})

	// Start the Gin server
	fmt.Println("Server is running on :8080")
	r.Run(":8080")
}
