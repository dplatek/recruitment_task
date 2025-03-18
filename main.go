package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"sort"

	"github.com/gin-gonic/gin"
)

func loadDataFromFile(filename string) ([]int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
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
			return nil, fmt.Errorf("invalid number in file: %s", line)
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

func loadData() ([]int, error) {
	fileData, err := loadDataFromFile("input.txt")
	if err != nil {
		return nil, err
	}

	sort.Ints(fileData)

	return fileData, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findCloseEnoughValue(data []int, value int, margin float64) (int, int) {
	left, right := 0, len(data)-1

	for left <= right {
		mid := (left + right) / 2
		diff := float64(abs(data[mid] - value))

		if diff <= margin {
			return data[mid], mid
		}

		if data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return 0, -1
}

func endpointHandler(c *gin.Context, data []int) {
	valueStr := c.Param("value")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value, must be an integer", "index": -1})
		return
	}

	margin := float64(value) * 0.10

	closestValue, closestIndex := findCloseEnoughValue(data, value, margin)

	if closestValue == value {
		c.JSON(http.StatusOK, gin.H{"error": "", "index": closestIndex})
		return
	}

	if closestValue != 0 {
		c.JSON(http.StatusOK, gin.H{
			"error": fmt.Sprintf("Value %d not found, but closest match %d found at index %d", value, closestValue, closestIndex),
			"index": closestIndex,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": fmt.Sprintf("Value %d not found", value),
		"index": -1,
	})
}

func main() {
	data, err := loadData()
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}

	r := gin.Default()
	r.GET("/endpoint/:value", func(c *gin.Context) {
		endpointHandler(c, data)
	})

	fmt.Println("Server is running on :8080")
	r.Run(":8080")
}
