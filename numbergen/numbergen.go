package main

import "time"
import "fmt"
import "math/rand"

func main() {
	randsrc := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randsrc)
	fmt.Print(rnd.Intn(100))
}
