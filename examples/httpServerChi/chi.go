package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	// "github.com/go-chi/render"
)

func main() {
	router := chi.NewRouter()

	router.Get("/thing/{uuid}", GetThing)
	//router.Post("/thing/{uuid}", UpdateThing)
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
	uuid := chi.URLParam(r, "uuid")

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

	//render.JSON(w, r, thingResponse)
}
