package db

import (
	"errors"
	"time"
)

var (
	ErrThingNotFound = errors.New("thing not found")
)

type DB interface {
	GetThing(uuid string) (Thing, error)
	CreateThing(name string, value string) (Thing, error)
	UpdateThing(uuid string, value string) (Thing, error)
	DeleteThing(uuid string) error
	ListThings(offset int, limit int) ([]Thing, int, error)
	GetSize() (int, error)
}

type Thing struct {
	UUID    string
	Name    string
	Value   string
	Updated time.Time
	Created time.Time
}
