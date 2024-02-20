package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello Mod")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey Hello from MOD")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Server running using Golang</h1>"))
}
