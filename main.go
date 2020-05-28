package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	facts []string = []string{"Dublin, Ireland", "London, United Kingdom"}
)

func main() {
	fmt.Println("Factitude!")
	rand.Seed(time.Now().UnixNano())
	randomN := rand.Intn(len(facts))
	fmt.Println(facts[randomN])
	fmt.Println("\u2588\u2588\u2588")
}
