package main

import (
	"encoding/json"
	"log"
	"net/http"

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
	s.router.HandleFunc("/thing/{uuid}", s.GetThing)
}

func (s *server) ListenAndServe(addr string) {
	hs := &http.Server{Addr: addr, Handler: s.router}
	if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

type ThingResponse struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Value string `json:"value"`
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

	thingResponse := ThingResponse{
		UUID:  thing.UUID,
		Name:  thing.Name,
		Value: thing.Value,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(thingResponse)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
