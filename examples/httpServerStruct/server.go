package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ldej/go-training/examples/db"
	"github.com/ldej/go-training/examples/db/inmemory"
)

func main() {
	dbService := inmemory.NewDB()
	server := NewServer(dbService)
	server.ListenAndServe(":8080")
}

type server struct {
	router *mux.Router
	db     db.DB
}

func NewServer(db db.DB) *server {
	s := &server{
		db: db,
	}
	s.Routes()
	return s
}

func (s *server) Routes() {
	s.router = mux.NewRouter()
	s.router.HandleFunc("/thing/new", s.CreateThing)
	s.router.HandleFunc("/thing/{uuid}", s.GetThing)
	s.router.HandleFunc("/thing/{uuid}", s.UpdateThing)
	s.router.HandleFunc("/thing/{uuid}", s.DeleteThing)
}

func (s *server) ListenAndServe(addr string) {
	hs := &http.Server{Addr: addr, Handler: s.router}
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

type ThingResponse struct {
	UUID    string    `json:"uuid"`
	Name    string    `json:"name"`
	Value   string    `json:"value"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

func (s *server) GetThing(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	thing, err := s.db.GetThing(uuid)
	if err == db.ErrThingNotFound {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	thingResponse := thingToThingResponse(thing)

	JSON(w, thingResponse)
}

type CreateThing struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (s *server) CreateThing(w http.ResponseWriter, r *http.Request) {
	var createThing CreateThing
	err := json.NewDecoder(r.Body).Decode(&createThing)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	createdThing, err := s.db.CreateThing(createThing.Name, createThing.Value)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	thingResponse := thingToThingResponse(createdThing)

	JSON(w, thingResponse)
}

type UpdateThing struct {
	Value string `json:"value"`
}

func (s *server) UpdateThing(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	var updateThing UpdateThing
	err := json.NewDecoder(r.Body).Decode(&updateThing)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	updatedThing, err := s.db.UpdateThing(uuid, updateThing.Value)
	if err == db.ErrThingNotFound {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	thingResponse := thingToThingResponse(updatedThing)

	JSON(w, thingResponse)
}

func (s *server) DeleteThing(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	err := s.db.DeleteThing(uuid)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type ThingsResponse struct {
	Total  int             `json:"total"`
	Page   int             `json:"page"`
	Limit  int             `json:"limit"`
	Things []ThingResponse `json:"things"`
}

func (s *server) ListThings(w http.ResponseWriter, r *http.Request) {
	page := 1
	pageQueryParam := r.URL.Query().Get("page")
	if pageQueryParam != "" {
		result, err := strconv.Atoi(pageQueryParam)
		if err == nil && result > 0 {
			page = result
		}
	}

	limit := 10
	limitQueryParam := r.URL.Query().Get("limit")
	if limitQueryParam != "" {
		result, err := strconv.Atoi(limitQueryParam)
		if err == nil && result > 0 && result <= 100 {
			limit = result
		}
	}

	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}

	things, count, err := s.db.ListThings(offset, limit)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	thingsResponse := ThingsResponse{
		Page:   page,
		Limit:  limit,
		Total:  count,
		Things: []ThingResponse{},
	}
	for _, thing := range things {
		thingsResponse.Things = append(thingsResponse.Things, thingToThingResponse(thing))
	}
	JSON(w, thingsResponse)
}

func JSON(w http.ResponseWriter, a interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(a)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func thingToThingResponse(thing db.Thing) ThingResponse {
	return ThingResponse{
		UUID:    thing.UUID,
		Name:    thing.Name,
		Value:   thing.Value,
		Updated: thing.Updated,
		Created: thing.Created,
	}
}

// UpdateThing
// CreateThing
// DeleteThing
