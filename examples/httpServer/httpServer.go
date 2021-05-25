package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/thing/", GetThing)
	// For Windows users: use localhost:8080 instead of :8080 to prevent a security message
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err)
	}
}

type ThingResponse struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Value string `json:"value"`

	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

func GetThing(w http.ResponseWriter, r *http.Request) {
	// TODO Which method? GET, POST, PUT?

	result := strings.Split(r.URL.Path, "/")
	uuid := result[len(result)-1]

	thingResponse := ThingResponse{
		UUID:  uuid,
		Name:  "example",
		Value: "example",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(thingResponse)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//w.WriteHeader(http.StatusOK) Not needed, as Encode uses w.Write which sets the status header to OK
}
