package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/thing/{uuid}", GetThing).Methods(http.MethodGet)
	//router.HandleFunc("/thing/{uuid}", UpdateThing).Methods(http.MethodPost)

	if err := http.ListenAndServe(":8080", router); err != nil {
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
	uuid := mux.Vars(r)["uuid"]

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
}
