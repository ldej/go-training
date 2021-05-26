package sqlite3

import (
	"os"
	"testing"

	"github.com/ldej/go-training/examples/db"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	db db.DB
}

func (s *Suite) SetupSuite() {
	testFilename := "./test-sqlite3.db"
	_ = os.Remove(testFilename)
	var err error
	s.db, err = NewDB(testFilename)
	s.NoError(err)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestThing() {
	createdThing, err := s.db.CreateThing("name", "value")
	s.NoError(err)

	retrievedThing, err := s.db.GetThing(createdThing.UUID)
	s.NoError(err)
	s.Equal(createdThing.Value, retrievedThing.Value)

	updatedThing, err := s.db.UpdateThing(retrievedThing.UUID, "updated")
	s.NoError(err)
	s.Equal("updated", updatedThing.Value)

	things, count, err := s.db.ListThings(0, 100)
	s.NoError(err)
	s.Equal(1, count)
	s.Equal(len(things), count)

	err = s.db.DeleteThing(retrievedThing.UUID)
	s.NoError(err)
}

func (s *Suite) TestThingNotFound() {
	_, err := s.db.GetThing("does-not-exist")
	s.Equal(db.ErrThingNotFound, err)

	_, err = s.db.UpdateThing("does-not-exist", "")
	s.Equal(db.ErrThingNotFound, err)
}
