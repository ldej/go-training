package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	// mock response
	recorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "/thing/123", nil)
	assert.NoError(t, err)
	request.Header.Set("Accept", "application/json")

	GetThing(recorder, request)

	//  verify response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// decode json
	var thing ThingResponse
	err = json.NewDecoder(recorder.Body).Decode(&thing)
	assert.NoError(t, err)

	assert.Equal(t, "123", thing.UUID)
}
