package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"recruitment_task/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEndpointHandler(t *testing.T) {
	data := []int{0, 10, 20, 100, 1000000}

	r := gin.Default()
	r.GET("/endpoint/:value", func(c *gin.Context) {
		handlers.EndpointHandler(c, data)
	})

	tests := []struct {
		name        string
		value       string
		expectCode  int
		expectErr   string
		expectIndex int
	}{
		{"Valid value - exact match", "10", http.StatusOK, "", 1},
		{"Valid value - closest match", "21", http.StatusOK, "Value 21 not found, but closest match 20 found at index 2", 2},
		{"Value not found", "200", http.StatusNotFound, "Value 200 not found", -1},
		{"Invalid value", "abc", http.StatusBadRequest, "Invalid value, must be an integer", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/endpoint/"+tt.value, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectCode, w.Code)

			expectedJSON := fmt.Sprintf(`{"error": "%s", "index": %d}`, tt.expectErr, tt.expectIndex)
			assert.JSONEq(t, expectedJSON, w.Body.String())
		})
	}
}
