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
		fmt.Fprintf(w, "step 2")
		log.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
	}
	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
