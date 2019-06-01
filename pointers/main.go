package main

import "fmt"

func main() {
	i := 1

	//unlike c, Go does not have pointer arithmetic
	var p *int

	p = &i

	fmt.Println(p)
	*p++

	fmt.Println(i)
}
