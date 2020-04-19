package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/stretchr/testify/assert"
)

func TestGetThings(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Init", 301, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupRouter()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/Viz", nil)

			router.ServeHTTP(w, req)
			assert.Equal(t, tt.want, w.Code)
		})
	}
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
