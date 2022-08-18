package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Status struct {
	Message string `json:"message"`
	Status  string `json: "status"`
}

func pingPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := Status{
		Message: "All is good with the world",
		Status:  "Success",
	}
	json.NewEncoder(w).Encode(message)
}

func handleRequests() {
	http.HandleFunc("/ping", pingPage)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	handleRequests()
}
