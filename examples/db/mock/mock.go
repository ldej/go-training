package mock

import "github.com/ldej/go-training/examples/db"

var _ db.DB = (*mock)(nil)

/* A mock for the db struct

Inspired by https://github.com/benbjohnson/wtf/blob/main/mock/dial.go

With this mock, you can replace the function with your desired function within your test
*/

type mock struct {
	GetThingFn    func(uuid string) (db.Thing, error)
	CreateThingFn func(name string, value string) (db.Thing, error)
	UpdateThingFn func(uuid string, value string) (db.Thing, error)
	DeleteThingFn func(uuid string) error
	ListThingsFn  func(offset int, limit int) ([]db.Thing, int, error)
	GetSizeFn     func() (int, error)
}

func (m *mock) GetThing(uuid string) (db.Thing, error) {
	return m.GetThingFn(uuid)
}

func (m *mock) CreateThing(name string, value string) (db.Thing, error) {
	return m.CreateThingFn(name, value)
}

func (m *mock) UpdateThing(uuid string, value string) (db.Thing, error) {
	return m.UpdateThingFn(uuid, value)
}

func (m *mock) DeleteThing(uuid string) error {
	return m.DeleteThingFn(uuid)
}

func (m *mock) ListThings(offset int, limit int) ([]db.Thing, int, error) {
	return m.ListThingsFn(offset, limit)
}

func (m *mock) GetSize() (int, error) {
	return m.GetSizeFn()
}
