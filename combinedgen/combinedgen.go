package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var namegenUrl string = ""
var numbergenUrl string = ""

func main() {
	http.HandleFunc("/", handler)
	setServiceUrls()
	log.Print("Listening.\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Handling request.\n")
	nameResponse, err := http.Get(namegenUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer nameResponse.Body.Close()
	nameBuf, err := ioutil.ReadAll(nameResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	name := string(nameBuf)

	numberResponse, err := http.Get(numbergenUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer numberResponse.Body.Close()
	numberBuf, err := ioutil.ReadAll(numberResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	number := string(numberBuf)

	name = strings.TrimSpace(name)
	fmt.Fprintf(w, "%s-%s", name, number)
}

func setServiceUrls() {
	namegenHost := os.Getenv("NAMEGENSERVICE_SERVICE_HOST")
	log.Printf("namegen host: %s", namegenHost)
	namegenPort := os.Getenv("NAMEGENSERVICE_SERVICE_PORT")
	log.Printf("namegen port: %s", namegenPort)
	namegenUrl = fmt.Sprintf("http://%s:%s", namegenHost, namegenPort)
	log.Printf("namegen service url: %s", namegenUrl)

	numbergenHost := os.Getenv("NUMBERGENSERVICE_SERVICE_HOST")
	log.Printf("numbergen host: %s", numbergenHost)
	numbergenPort := os.Getenv("NUMBERGENSERVICE_SERVICE_PORT")
	log.Printf("numbergen port: %s", numbergenPort)
	numbergenUrl = fmt.Sprintf("http://%s:%s", numbergenHost, numbergenPort)
	log.Printf("numbergen service url: %s", numbergenUrl)
}
