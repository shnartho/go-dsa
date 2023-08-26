/*
Problem: Simple HTTP Server
Write a Go program that sets up a simple HTTP server. When a user accesses the server via a web browser or using tools like curl, the server should respond with a "Hello, World!" message.
*/

package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8800", nil)
}
