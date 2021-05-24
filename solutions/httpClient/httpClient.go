package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type GetThingResponse struct {
	UUID    string    `json:"uuid"`
	Name    string    `json:"name"`
	Value   string    `json:"value"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

type Client struct {
	Hostname string
}

func (cl *Client) GetThingOnUUID(thingUUID string) (*GetThingResponse, error) {
	client := http.Client{}
	httpResponse, err := client.Get(fmt.Sprintf("%s/thing/%s", cl.Hostname, thingUUID))
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching thing: http-status %d", httpResponse.StatusCode)
	}
	var resp GetThingResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func main() {
	client := Client{Hostname: "https://api-ldej-nl.el.r.appspot.com"}
	resp, err := client.GetThingOnUUID("4ff83452-5878-11ea-bc7e-914aa98404f8")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%v", resp)
}
