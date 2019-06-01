package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3, 4)
	fmt.Println(s1, s2)

	s1[0] = 999
	fmt.Println(s1, s2)
}
