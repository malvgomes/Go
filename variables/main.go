package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	//single-line declarations
	var i, j, k int
	var b, f, s = true, 2.3, "four"

	//multiple return function declaration
	var g, err = os.Open("name")
	fmt.Println(i, j, k, b, f, s, g, err)

	//short variable declaration
	freq := rand.Float64() * 3.0
	t := 0.0
	fmt.Println(freq, t)

	//multiple initialization
	a, c := 0, 1

	//quick swap
	a, c = c, a
	fmt.Println(a, c)
}
