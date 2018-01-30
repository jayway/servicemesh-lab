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

var failToggle = false

func main() {
	http.HandleFunc("/", handler)
	log.Print("Listening.\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Handling request.\n")
	envFail := os.Getenv("GENERATE_ERRORS")
	envFail = strings.ToLower(envFail)
	log.Printf("Fail set to: %s\n", envFail)
	if envFail == "true" && failToggle {
		http.Error(w, "Toggled error", http.StatusInternalServerError)
	} else {
		randsrc := rand.NewSource(time.Now().UnixNano())
		rnd := rand.New(randsrc)
		fmt.Fprintf(w, "%d\n", rnd.Intn(100))
	}
	failToggle = !failToggle
}
