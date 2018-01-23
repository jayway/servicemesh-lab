package main

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"log"
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	randsrc := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randsrc)
	name := animals[rnd.Intn(len(animals))]
	fmt.Fprintf(w, "%s\n", name)
}
