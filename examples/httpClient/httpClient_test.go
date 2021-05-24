package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThingByUUIDSuccess(t *testing.T) {
	// create server simulator
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/thing/123", r.URL.RequestURI())

		returnSuccessResponse(t, w, r)
	}))
	defer testServer.Close()

	// perform request
	client := &Client{Hostname: testServer.URL}
	resp, err := client.GetThingOnUUID("123")

	// validate output as return from fake server
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Laurence", resp.Name)
	assert.Equal(t, "cheese", resp.Value)
}

func returnSuccessResponse(t *testing.T, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := Thing{
		Name:  "Laurence",
		Value: "cheese",
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		t.Fatalf("Error encoding response")
	}
}

func TestGetThingByUUIDError(t *testing.T) {
	// create server simulator
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer testServer.Close()

	// perform request
	client := &Client{Hostname: testServer.URL}
	_, err := client.GetThingOnUUID("123")

	// validate output as return from fake server
	assert.Error(t, err)
}
