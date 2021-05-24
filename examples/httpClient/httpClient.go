package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
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
	Hostname   string
	HTTPClient http.Client
}

func NewThingClient(hostname string) *Client {
	return &Client{
		Hostname:   hostname,
		HTTPClient: http.Client{Timeout: 30 * time.Second},
	}
}

func (cl *Client) GetThingOnUUID(thingUUID string) (Thing, error) {
	httpResponse, err := cl.HTTPClient.Get(fmt.Sprintf("%s/thing/%s", cl.Hostname, thingUUID))
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
	thingToCreate := CreateThing{Name: name, Value: value}

	thingJSON, err := json.Marshal(thingToCreate)
	if err != nil {
		return Thing{}, err
	}
	httpResponse, err := cl.HTTPClient.Post(fmt.Sprintf("%s/thing/new", cl.Hostname), "application/json", bytes.NewBuffer(thingJSON))
	if err != nil {
		return Thing{}, err
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != http.StatusOK {
		return Thing{}, fmt.Errorf("error creating thing: http-status %d", httpResponse.StatusCode)
	}

	var createdThing Thing
	err = json.NewDecoder(httpResponse.Body).Decode(&createdThing)
	if err != nil {
		return Thing{}, err
	}
	return createdThing, nil
}

type UpdateThing struct {
	Value string `json:"value"`
}

func (cl *Client) UpdateThing(uuid string, value string) (Thing, error) {
	thingToUpdate := UpdateThing{Value: value}

	thingJSON, err := json.Marshal(thingToUpdate)
	if err != nil {
		return Thing{}, err
	}
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/thing/%s", cl.Hostname, uuid), bytes.NewBuffer(thingJSON))
	if err != nil {
		return Thing{}, err
	}

	httpResponse, err := cl.HTTPClient.Do(request)
	if err != nil {
		return Thing{}, err
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode != http.StatusOK {
		return Thing{}, fmt.Errorf("error updating thing: http-status %d", httpResponse.StatusCode)
	}

	var updatedThing Thing
	err = json.NewDecoder(httpResponse.Body).Decode(&updatedThing)
	if err != nil {
		return Thing{}, err
	}
	return updatedThing, nil
}

func (cl *Client) DeleteThing(uuid string) error {
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/thing/%s", cl.Hostname, uuid), nil)
	if err != nil {
		return err
	}

	httpResponse, err := cl.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	if httpResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("error deleting thing: http-status %d", httpResponse.StatusCode)
	}
	return nil
}

type ThingsList struct {
	Page   int     `json:"page"`
	Total  int     `json:"total"`
	Limit  int     `json:"limit"`
	Things []Thing `json:"things"`
}

func (cl *Client) ListThings(page int, limit int) (ThingsList, error) {
	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/thing", cl.Hostname), nil)
	if err != nil {
		return ThingsList{}, err
	}
	queryParams := url.Values{}
	if page > 0 {
		queryParams.Set("page", strconv.Itoa(page))
	}
	if limit > 0 {
		queryParams.Set("limit", strconv.Itoa(limit))
	}
	request.URL.RawQuery = queryParams.Encode()

	httpResponse, err := client.Do(request)
	if err != nil {
		return ThingsList{}, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return ThingsList{}, fmt.Errorf("error fetching things: http-status %d", httpResponse.StatusCode)
	}
	var thingsList ThingsList
	err = json.NewDecoder(httpResponse.Body).Decode(&thingsList)
	if err != nil {
		return ThingsList{}, err
	}
	return thingsList, nil
}

func main() {
	client := NewThingClient("https://api-ldej-nl.el.r.appspot.com")

	createdThing, err := client.CreateThing("Laurence", "hello")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created Thing: %v\n", createdThing)

	thing, err := client.GetThingOnUUID(createdThing.UUID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ReceivedThing: %v\n", thing)

	updatedThing, err := client.UpdateThing(thing.UUID, "updated")
	fmt.Printf("UpdatedThing: %v\n", updatedThing)

	thingsList, err := client.ListThings(1, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Things List: %v\n", thingsList)

	err = client.DeleteThing(thing.UUID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully deleted thing: %v\n", thing)
}
