package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEndpointHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	input := []int{0, 10, 20, 30, 40, 100, 1000000}
	r := gin.Default()
	r.GET("/endpoint/:value", func(c *gin.Context) {
		EndpointHandler(c, input)
	})

	tests := []struct {
		name        string
		value       string
		expectCode  int
		expectErr   string
		expectIndex int
	}{
		{"Valid exact match", "10", http.StatusOK, "", 1},
		{"Closest match", "21", http.StatusOK, "Value 21 not found, but closest match 20 found at index 2", 2},
		{"Value not found", "1000", http.StatusNotFound, "Value 1000 not found", -1},
		{"Invalid input", "abc", http.StatusBadRequest, "Invalid value, must be an integer", -1},
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
