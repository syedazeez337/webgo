package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", home)
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/greeting/{morning}", morning)
	mux.HandleFunc("/greeting/good/{night}", night)

	log.Print("Server starting...")

	err := http.ListenAndServe(":8080", mux)
	log.Panic(err)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is home!!"))
}

func morning(w http.ResponseWriter, r *http.Request) {
	/*
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}
	*/

	id := r.PathValue("morning")

	msg := fmt.Sprintf("Good %s from the server\n", id)
	w.Write([]byte(msg))
}

func night(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("night")

	msg := fmt.Sprintf("Good %s from the server\n", id)
	w.Write([]byte(msg))
}
