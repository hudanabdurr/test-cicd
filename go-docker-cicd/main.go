package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	// Get port from environment variable or set a default
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080" // Default port if APP_PORT is not set
	}

	log.Printf("Server started at http://localhost:%s", appPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", appPort), r))
}
