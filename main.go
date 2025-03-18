package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port     string `json:"port"`
	LogLevel string `json:"log_level"`
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func setLogLevel(level string) {
	switch level {
	case "Debug":
		log.SetOutput(os.Stdout)
	case "Info":
		log.SetOutput(os.Stdout)
	case "Error":
		log.SetOutput(os.Stderr)
	default:
		log.Printf("Unknown log level: %s. Defaulting to Info.", level)
		log.SetOutput(os.Stdout)
	}
}

func loadDataFromFile(filename string) ([]int, error) {
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
			return nil, fmt.Errorf("invalid number in file: %s", line)
		}
		numbers = append(numbers, number)
	}

	log.Printf("Info: Successfully loaded %d records from file", len(numbers))
	return numbers, nil
}

func findCloseEnoughValue(data []int, value int, margin float64) (int, int) {
	left, right := 0, len(data)-1
	var closestValue int
	var closestIndex int
	minDiff := margin + 1

	for left <= right {
		mid := left + (right-left)/2
		diff := float64(abs(data[mid] - value))

		if diff <= margin {
			if diff < minDiff {
				minDiff = diff
				closestValue = data[mid]
				closestIndex = mid
			}
			if data[mid] < value {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return closestValue, closestIndex
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func endpointHandler(c *gin.Context, data []int) {
	valueStr := c.Param("value")
	log.Printf("Info: Received request: %s %s with value %s", c.Request.Method, c.Request.URL.Path, valueStr)

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Error: Invalid value provided: %s", valueStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value, must be an integer", "index": -1})
		return
	}

	log.Printf("Debug: Processing value: %d", value)

	margin := float64(value) * 0.10
	closestValue, closestIndex := findCloseEnoughValue(data, value, margin)

	if value == closestValue {
		log.Printf("Info: Exact match found for value %d at index %d", value, closestIndex)
		c.JSON(http.StatusOK, gin.H{"error": "", "index": closestIndex})
		return
	}

	if closestValue != 0 {
		log.Printf("Info: Closest match found for value %d at index %d: %d", value, closestIndex, closestValue)
		c.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("Value %d not found, but closest match %d found at index %d", value, closestValue, closestIndex), "index": closestIndex})
		return
	}

	log.Printf("Error: Value %d not found, sending 404 response", value)
	c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Value %d not found", value), "index": -1})
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	setLogLevel(config.LogLevel)

	log.Printf("Info: Server is starting on port %s", config.Port)

	data, err := loadDataFromFile("input.txt")
	if err != nil {
		log.Fatalf("Error loading data: %v", err)
		return
	}

	r := gin.Default()
	r.GET("/endpoint/:value", func(c *gin.Context) {
		endpointHandler(c, data)
	})

	log.Printf("Info: Server is running on :%s", config.Port)
	r.Run(":" + config.Port)
}
