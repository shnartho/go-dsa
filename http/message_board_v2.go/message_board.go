/*
A simple complete mwessage board api rest project
*/

package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
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

	response, _ := json.Marshal(messages)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	messages = append(messages, message.Text)
	w.WriteHeader(http.StatusCreated)
}

func updateMessageHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	vars := mux.Vars(r)
	indexStr := vars["index"]
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(messages) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}

	var message Message
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}

	messages[index] = message.Text
	w.WriteHeader(http.StatusOK)
}

func deleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	vars := mux.Vars(r)
	indexStr := vars["index"]
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(messages) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}
	messages = append(messages[:index], messages[index+1:]...)
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/messages", getMessagesHandler).Methods(http.MethodGet)
	router.HandleFunc("/messages", postMessageHandler).Methods(http.MethodPost)
	router.HandleFunc("/messages/{index:[0-9]+}", updateMessageHandler).Methods(http.MethodPut)    // [0-9]+ reg expresion
	router.HandleFunc("/messages/{index:[0-9]+}", deleteMessageHandler).Methods(http.MethodDelete) //[0-9]+ one/more digits

	http.Handle("/", router) // "/" means router will handle/match all endpoints after /

	http.ListenAndServe(":8800", nil)

}
