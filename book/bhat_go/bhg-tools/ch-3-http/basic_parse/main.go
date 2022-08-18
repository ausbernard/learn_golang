package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Status struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	res, err := http.Post(
		"http://localhost:5000/ping",
		"application/json",
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}

	var status Status
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	log.Printf("%s -> %s\n", status.Status, status.Message)

}
