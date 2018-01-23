package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Call namegen
	name := "albatross"
	// Call numbergen
	number := "42"

	fmt.Fprintf(w, "%s-%s", name, number)
}
