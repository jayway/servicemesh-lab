package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	animals = [...]string{
		"aardvark",
		"albatross",
		"alligator",
		"Alpaca",
		"ant",
		"anteater",
		"antelope",
		"ape",
		"armadillo",
	}
)

func main() {
	randsrc := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randsrc)
	name := animals[rnd.Intn(len(animals))]
	fmt.Print(name)
}
