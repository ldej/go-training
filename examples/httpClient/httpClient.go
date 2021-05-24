package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Thing struct {
	UUID    string    `json:"uuid"`
	Name    string    `json:"name"`
	Value   string    `json:"value"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

type Client struct {
	Hostname string
}

func (cl *Client) GetThingOnUUID(thingUUID string) (Thing, error) {
	client := http.Client{}
	httpResponse, err := client.Get(fmt.Sprintf("%s/thing/%s", cl.Hostname, thingUUID))
	if err != nil {
		return Thing{}, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return Thing{}, fmt.Errorf("error fetching thing: http-status %d", httpResponse.StatusCode)
	}
	var thing Thing
	err = json.NewDecoder(httpResponse.Body).Decode(&thing)
	if err != nil {
		return Thing{}, err
	}
	return thing, nil
}

type CreateThing struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (cl *Client) CreateThing(name string, value string) (Thing, error) {
	thing := CreateThing{Name: name, Value: value}

	client := http.Client{}
	thingJSON, err := json.Marshal(thing)
	if err != nil {
		return Thing{}, err
	}
	httpResponse, err := client.Post(fmt.Sprintf("%s/thing/new", cl.Hostname), "application/json", bytes.NewBuffer(thingJSON))
	if err != nil {
		return Thing{}, err
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != http.StatusOK {
		return Thing{}, errors.New("not ok")
	}

	var resp Thing
	err = json.NewDecoder(httpResponse.Body).Decode(&resp)
	if err != nil {
		return Thing{}, err
	}
	return resp, nil
}

type UpdateThing struct {
	Value string `json:"value"`
}

func (cl *Client) UpdateThing(uuid string, value string) (Thing, error) {
	client := http.Client{}

	thing := UpdateThing{Value: value}
	thingJSON, err := json.Marshal(thing)
	if err != nil {
		return Thing{}, err
	}
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/thing/%s", cl.Hostname, uuid), bytes.NewBuffer(thingJSON))
	if err != nil {
		return Thing{}, err
	}

	httpResponse, err := client.Do(request)
	_ = httpResponse // TODO for you to finish
	return Thing{}, nil
}

// TODO DELETE
// TODO List

func main() {
	client := Client{Hostname: "https://api-ldej-nl.el.r.appspot.com"}

	thing, err := client.CreateThing("Laurence", "hello")
	if err != nil {
		log.Fatal(err)
	}

	thing, err = client.GetThingOnUUID(thing.UUID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%v", thing)
}
