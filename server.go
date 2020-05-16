package main

import "fmt"
import "log"
import "net/http"

func main() {
	fmt.Println("vim-go")

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "henlo")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8111", nil))
}
