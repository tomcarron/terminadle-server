// write a go web server which provides an api for a 5 letter word of the day.
// The word of the day should be randomly selected from a list of 5 letter words.
// The server should have a single endpoint which returns the word of the day.
// The server should also have a healthcheck endpoint which returns a 200 status code.
// The server should have tests which test the word of the day endpoint and the healthcheck endpoint.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var words = []string{"apple", "audio", "mango", "grape", "melon"}

// WordOfTheDay struct
type WordOfTheDay struct {
	Word string `json:"word"`
}

// GetWordOfTheDay function
func GetWordOfTheDay(w http.ResponseWriter, r *http.Request) {
	word := words[rand.Intn(len(words))]
	wordOfTheDay := WordOfTheDay{Word: word}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wordOfTheDay)
}

// HealthCheck function
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/word-of-the-day", GetWordOfTheDay)
	http.HandleFunc("/healthcheck", HealthCheck)
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
