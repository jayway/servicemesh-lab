package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	animals = [...]string{
		"aardvark",
		"albatross",
		"alligator",
		"alpaca",
		"ant",
		"anteater",
		"antelope",
		"ape",
		"armadillo",
	}
)
var failToggle = false
var generateErrors = false

func main() {
	http.HandleFunc("/", handler)
	generateErrorsEnv := os.Getenv("GENERATE_ERRORS")
	generateErrorsEnv = strings.ToLower(generateErrorsEnv)
	if generateErrorsEnv == "true" {
		generateErrors = true
	}
	log.Printf("Generate errors: %t\n", generateErrors)
	log.Print("Listening.\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Handling request.\n")
	if generateErrors && failToggle {
		log.Print("Failing!\n")
		http.Error(w, "Toggled error", http.StatusInternalServerError)
	} else {
		randsrc := rand.NewSource(time.Now().UnixNano())
		rnd := rand.New(randsrc)
		name := animals[rnd.Intn(len(animals))]
		log.Print("Success.\n")
		fmt.Fprintf(w, "%s\n", name)
	}
	failToggle = !failToggle
}
