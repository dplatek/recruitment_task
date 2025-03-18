package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEndpointHandler(t *testing.T) {
	// Example data that will be loaded from the input.txt file
	data := []int{0, 10, 20, 100, 1000000}

	// Initialize the Gin router
	r := gin.Default()

	// Define the route with the handler
	r.GET("/endpoint/:value", func(c *gin.Context) {
		endpointHandler(c, data)
	})

	tests := []struct {
		name       string
		value      string
		expectCode int
		expectBody string
	}{
		{"Valid value", "10", http.StatusOK, "Value 10 found at index 1"},
		{"Value not found", "200", http.StatusNotFound, "Value 200 not found"},
		{"Invalid value", "abc", http.StatusBadRequest, "Invalid value, must be an integer"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request using value instead of index
			req := httptest.NewRequest("GET", "/endpoint/"+tt.value, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Assert status code and body
			assert.Equal(t, tt.expectCode, w.Code)
			assert.Equal(t, tt.expectBody, w.Body.String())
		})
	}
}
