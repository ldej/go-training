package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestThingByUUIDSuccess(t *testing.T) {
	// create server simulator
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/thing/123", r.URL.RequestURI())

		returnSuccessResponse(t, w, r)
	}))
	defer testServer.Close()

	// perform request
	client := NewThingClient(testServer.URL)
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

	fixedDate := time.Date(2021, 5, 24, 0, 0, 0, 0, time.UTC)

	resp := Thing{
		UUID:    "123",
		Name:    "Laurence",
		Value:   "cheese",
		Updated: fixedDate,
		Created: fixedDate,
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
	client := NewThingClient(testServer.URL)
	_, err := client.GetThingOnUUID("123")

	// validate output as return from fake server
	assert.Error(t, err)
}

func TestCreateThing(t *testing.T) {
	// create server simulator
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/thing/new", r.URL.RequestURI())

		returnSuccessResponse(t, w, r)
	}))
	defer testServer.Close()

	// perform request
	client := NewThingClient(testServer.URL)
	createdThing, err := client.CreateThing("Laurence", "cheese")

	// validate output as return from fake server
	assert.NoError(t, err)
	assert.Equal(t, "Laurence", createdThing.Name)
	assert.Equal(t, "cheese", createdThing.Value)
}

func TestUpdateThing(t *testing.T) {
	// create server simulator
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/thing/123", r.URL.RequestURI())

		returnSuccessResponse(t, w, r)
	}))
	defer testServer.Close()

	// perform request
	client := NewThingClient(testServer.URL)
	createdThing, err := client.UpdateThing("123", "cheese")

	// validate output as return from fake server
	assert.NoError(t, err)
	assert.Equal(t, "cheese", createdThing.Value)
}

func TestDeleteThing(t *testing.T) {
	// create server simulator
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/thing/123", r.URL.RequestURI())
	}))
	defer testServer.Close()

	// perform request
	client := NewThingClient(testServer.URL)
	err := client.DeleteThing("123")
	assert.NoError(t, err)
}

func TestListThings(t *testing.T) {
	// create server simulator
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/thing", r.URL.Path)
		assert.Equal(t, url.Values{"page": []string{"1"}, "limit": []string{"10"}}, r.URL.Query())

		returnListThingsSuccessResponse(t, w, r)
	}))
	defer testServer.Close()

	// perform request
	client := NewThingClient(testServer.URL)
	thingsList, err := client.ListThings(1, 10)
	assert.NoError(t, err)
	assert.Len(t, thingsList.Things, 1)
}

func returnListThingsSuccessResponse(t *testing.T, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	fixedDate := time.Date(2021, 5, 24, 0, 0, 0, 0, time.UTC)

	resp := ThingsList{
		Page:  1,
		Limit: 1,
		Total: 1,
		Things: []Thing{
			{
				UUID:    "123",
				Name:    "Laurence",
				Value:   "cheese",
				Updated: fixedDate,
				Created: fixedDate,
			},
		},
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		t.Fatalf("Error encoding response")
	}
}
