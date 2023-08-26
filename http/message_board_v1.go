/*
Problem: Basic Message Board API
Write a Go program that sets up a basic HTTP server for a message board API. The API should support two endpoints:

GET /messages: Retrieve all stored messages.
POST /messages: Add a new message to the message board.
*/

package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	messages []string
	mutex    sync.Mutex
)

type Message struct {
	Text string `json:"message"`
}

func getMessagesHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	respose, _ := json.Marshal(messages)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respose)
}

func postMessagesHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}

	messages = append(messages, message.Text)
	w.WriteHeader(http.StatusCreated)
}

func main() {
	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getMessagesHandler(w, r)
		case http.MethodPost:
			postMessagesHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8800", nil)
}
