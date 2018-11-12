package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have visited %s!", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handle)

	env := os.Getenv("APP_ENV")

	if env == "dev" {
		fmt.Println("Starting dev web server on port 8080")
	} else {
		fmt.Println("Starting prod web server on port 8080")
	}

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
