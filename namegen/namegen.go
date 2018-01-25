package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
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

func main() {
	http.HandleFunc("/", handler)
	log.Print("Listening.\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Handling request.\n")
	randsrc := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randsrc)
	name := animals[rnd.Intn(len(animals))]
	fmt.Fprintf(w, "%s\n", name)
}
