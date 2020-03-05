package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPatientByUIDSuccess(t *testing.T) {
	// create server simulator
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/api/v1/patients/123", r.URL.RequestURI())

		returnSuccessResponse(t, w, r)
	}))
	defer testServer.Close()

	// perform request
	client := &Client{Hostname: testServer.URL}
	resp, err := client.GetPatientOnUID("123")

	// validate output as return from fake server
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Laurence", resp.FullName)
	assert.Equal(t, []string{"cheese"}, resp.Allergies)
}

func returnSuccessResponse(t *testing.T, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := GetPatientResponse{
		FullName:  "Laurence",
		Allergies: []string{"cheese"},
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		t.Fatalf("Error encoding response")
	}
}

func TestGetPatientByUIDError(t *testing.T) {
	// create server simulator
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer testServer.Close()

	// perform request
	client := &Client{Hostname: testServer.URL}
	_, err := client.GetPatientOnUID("123")

	// validate output as return from fake server
	assert.Error(t, err)
}
