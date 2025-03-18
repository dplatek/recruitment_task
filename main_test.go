package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestEndpointHandler(t *testing.T) {
	r := gin.Default()
	r.GET("/endpoint/:index", endpointHandler)

	tests := []struct {
		name       string
		index      string
		expectCode int
		expectBody string
	}{
		{"Valid index", "100", http.StatusOK, "Success: You reached /endpoint/100"},
		{"Invalid index", "abc", http.StatusBadRequest, "Invalid index, must be an integer"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/endpoint/"+tt.index, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tt.expectCode {
				t.Errorf("expected status %d, got %d", tt.expectCode, w.Code)
			}
			if w.Body.String() != tt.expectBody {
				t.Errorf("expected body %q, got %q", tt.expectBody, w.Body.String())
			}
		})
	}
}
