package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raj5036/mongoapi/router"
)

func main() {
	/*
		MongoDB Credentials

		username: raj5036
		password: UEPzMoEAikdvu4ym
	*/
	fmt.Println("Using MongoDB with GO")
	fmt.Println("Server is getting started...........")

	router := router.Router()
	log.Fatal(http.ListenAndServe(":4000", router))
	fmt.Println("Listening on PORT: 4000...........")
}
