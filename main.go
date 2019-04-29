package main

import (
	"fmt"
	"net/http"

	mux "github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/greet/{name}", greeting)

	http.ListenAndServe(":8081", mux)
}

func greeting(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{"message": "Hello world %s"}`, name)
}
func homePage(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{"message": "Hello world"}`)
}
