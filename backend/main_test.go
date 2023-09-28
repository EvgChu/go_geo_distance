package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadRequestDistance(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/distance", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

}

func TestBadAddressDistance(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/distance?address=test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}
