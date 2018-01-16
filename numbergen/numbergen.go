package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randsrc := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randsrc)
	fmt.Print(rnd.Intn(100))
}
