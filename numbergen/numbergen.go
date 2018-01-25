package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
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
	fmt.Fprintf(w, "%d\n", rnd.Intn(100))
}
