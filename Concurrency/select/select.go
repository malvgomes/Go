package main

import (
	"fmt"
	"time"

	"github.com/malvgomes/html"
)

func fastest(url1, url2, url3 string) string {
	c1 := html.Title(url1)
	c2 := html.Title(url2)
	c3 := html.Title(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(180 * time.Millisecond):
		return "All lose"
		// default:
		// 	return "No answer"
	}
}

func main() {
	champ := fastest(
		"https://www.google.com",
		"https://www.cifraclub.com.br",
		"https://www.youtube.com",
	)

	fmt.Println(champ)
}
