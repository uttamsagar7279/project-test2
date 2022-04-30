package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	// Init the mux router
	router := mux.NewRouter()

	// Get all two data parameter from database
	router.HandleFunc("/movies/", GetEmp).Methods("GET")

	// Create a emp
	router.HandleFunc("/emp/", CreateEmp).Methods("POST")

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
