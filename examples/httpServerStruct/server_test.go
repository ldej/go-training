package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ldej/go-training/examples/db/inmemory"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	// mock response
	recorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "/thing/special", nil)
	assert.NoError(t, err)

	db := inmemory.NewDB()
	server := NewServer(db)
	go func() {
		server.ListenAndServe(":0")
	}()
	time.Sleep(100 * time.Millisecond)
	server.router.ServeHTTP(recorder, request)

	//  verify response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// decode json
	var thing ThingResponse
	err = json.NewDecoder(recorder.Body).Decode(&thing)
	assert.NoError(t, err)

	assert.Equal(t, "special", thing.UUID)
}
