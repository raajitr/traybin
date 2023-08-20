package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Hello": "World"})
}


func main() {
	fmt.Println("Started")
	r := mux.NewRouter()
	// s := r.PathPrefix("/products").Subrouter()
	r.HandleFunc("/", HomePage)
	http.ListenAndServe(":8080", r)
}