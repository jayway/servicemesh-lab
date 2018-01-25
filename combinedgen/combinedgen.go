package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var namegenurl string = ""
var numbergenurl string = ""

func main() {
	http.HandleFunc("/", handler)
	setserviceurls()
	log.Print("Listening.\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Handling request.\n")
	nameResponse, err := http.Get(namegenurl)
	if err != nil {
		log.Fatal(err)
	}
	defer nameResponse.Body.Close()
	nameBuf, err := ioutil.ReadAll(nameResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	name := string(nameBuf)

	numberResponse, err := http.Get(numbergenurl)
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

func setserviceurls() {
	namegenhost := os.Getenv("NAMEGENSERVICE_SERVICE_HOST")
	log.Printf("namegen host: %s", namegenhost)
	namegenport := os.Getenv("NAMEGENSERVICE_SERVICE_PORT")
	log.Printf("namegen port: %s", namegenport)
	namegenurl = fmt.Sprintf("http://%s:%s", namegenhost, namegenport)
	log.Printf("namegen service url: %s", namegenurl)

	numbergenhost := os.Getenv("NUMBERGENSERVICE_SERVICE_HOST")
	log.Printf("numbergen host: %s", numbergenhost)
	numbergenport := os.Getenv("NUMBERGENSERVICE_SERVICE_PORT")
	log.Printf("numbergen port: %s", numbergenport)
	numbergenurl = fmt.Sprintf("http://%s:%s", numbergenhost, numbergenport)
	log.Printf("numbergen service url: %s", numbergenurl)
}
