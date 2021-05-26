package main

import "net/http"

func main() {
	// http handler function
	// func(w http.ResponseWriter, r *http.Request)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Hello world"))
	})
	http.ListenAndServe("localhost:8080", nil)
}
