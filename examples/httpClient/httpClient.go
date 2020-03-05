package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GetPatientResponse struct {
	UID       string   `json:"uid"`
	FullName  string   `json:"full_name"`
	Allergies []string `json:"allergies"`
}

type Client struct {
	Hostname string
}

func (cl *Client) GetPatientOnUID(patientUid string) (*GetPatientResponse, error) {
	client := http.Client{}
	httpResponse, err := client.Get(fmt.Sprintf("%s/api/v1/patients/%s", cl.Hostname, patientUid))
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching patient: http-status %d", httpResponse.StatusCode)
	}
	var resp GetPatientResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func main() {
	client := Client{Hostname: "https://patient-store.appspot.com"}
	resp, err := client.GetPatientOnUID("4ff83452-5878-11ea-bc7e-914aa98404f8")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%v", resp)
}
