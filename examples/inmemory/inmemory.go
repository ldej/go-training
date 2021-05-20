package main

import (
	"errors"
	"fmt"
)

type Datastorer interface {
	Put(key string, value interface{}) error
	Get(key string) (interface{}, bool, error)
	Remove(key string) error
}

type InMemory struct {
	data map[string]interface{}
}

func NewInMemory() Datastorer {
	return &InMemory{
		data: map[string]interface{}{},
	}
}

func (i *InMemory) Get(key string) (interface{}, bool, error) {
	value, found := i.data[key]
	if !found {
		return nil, false, errors.New("not found")
	}
	return value, true, nil
}

func (i *InMemory) Put(key string, value interface{}) error {
	_, found := i.data[key]
	if found {
		return errors.New("key already exists")
	}
	i.data[key] = value
	return nil
}

func (i *InMemory) Remove(key string) error {
	_, found := i.data[key]
	if !found {
		return errors.New("key not found")
	}
	delete(i.data, key)
	return nil
}

func main() {
	inMemory := NewInMemory()
	err := inMemory.Put("laurence", "something")
	if err != nil {
		fmt.Println(err)
		return
	}
	result, found, err := inMemory.Get("laurence")
	fmt.Println(result, found, err)
}
