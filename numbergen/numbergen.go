package main

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	randsrc := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randsrc)
	fmt.Fprintf(w, "%s", rnd.Intn(100))
}
