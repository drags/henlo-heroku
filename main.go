package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "henlo")
	}
	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
