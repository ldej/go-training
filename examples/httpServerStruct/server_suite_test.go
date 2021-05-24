package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ldej/go-training/examples/db/inmemory"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	server *server
}

func (s *Suite) SetupSuite() {
	s.server = NewServer(inmemory.NewDB())
	go func() {
		s.server.ListenAndServe(":0")
	}()
	time.Sleep(10 * time.Millisecond)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) assertGetRoute(path string, statusCode int) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, path, nil)
	s.NoError(err)

	s.server.router.ServeHTTP(recorder, request)
	//  verify response
	s.Equal(statusCode, recorder.Code)
}

func (s *Suite) TestGetThing() {
	s.assertGetRoute("/thing/special", http.StatusOK)
}

func (s *Suite) TestGetThingNotFound() {
	s.assertGetRoute("/thing/does-not-exist", http.StatusNotFound)
}
