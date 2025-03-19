package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"recruitment_task/search"

	"github.com/gin-gonic/gin"
)

func EndpointHandler(c *gin.Context, input []int) {
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
	closestValue, closestIndex := search.FindCloseEnoughValue(input, value, margin)

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
