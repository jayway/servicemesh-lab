package main

import (
	"fmt"
	"log"
	"net/http"
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

	// To trim whitespace from the strings,
	// add "strings" to import and use strings.TrimSpace("gopher\n")
	fmt.Fprintf(w, "%s-%s\n", name, number)
}
