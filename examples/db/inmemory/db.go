package inmemory

import (
	"time"

	"github.com/ldej/go-training/examples/db"

	"github.com/google/uuid"
)

type service struct {
	things []db.Thing
}

func NewDB() *service {
	return &service{things: []db.Thing{{UUID: "special", Name: "Laurence", Value: "Coffee"}}}
}

func (s *service) GetThing(uuid string) (db.Thing, error) {
	for _, thing := range s.things {
		if thing.UUID == uuid {
			return thing, nil
		}
	}
	return db.Thing{}, db.ErrThingNotFound
}

func (s *service) CreateThing(name string, value string) (db.Thing, error) {
	now := time.Now().UTC()

	thing := db.Thing{
		UUID:    uuid.New().String(),
		Name:    name,
		Value:   value,
		Updated: now,
		Created: now,
	}

	s.things = append(s.things, thing)
	return thing, nil
}

func (s *service) UpdateThing(uuid string, value string) (db.Thing, error) {
	for i, thing := range s.things {
		if thing.UUID == uuid {
			thing.Value = value
			s.things[i] = thing
			return thing, nil
		}
	}
	return db.Thing{}, db.ErrThingNotFound
}

func (s *service) DeleteThing(uuid string) error {
	for i, thing := range s.things {
		if thing.UUID == uuid {
			s.things = append(s.things[:i], s.things[i+1:]...)
			return nil
		}
	}
	return db.ErrThingNotFound
}

func (s *service) ListThings(offset int, limit int) ([]db.Thing, int, error) {
	return s.things[offset : offset+limit], len(s.things), nil
}

func (s *service) GetSize() int {
	return len(s.things)
}
