package main

import (
	"fmt"

	"github.com/malvgomes/html"
)

func forward(origin <-chan string, destiny chan string) {
	for {
		destiny <- <-origin
	}
}

//multiplex group messages in a channel
func join(en1, en2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(en1, c)
	go forward(en2, c)

	return c
}

func main() {
	c := join(
		html.Title("https://www.cifraclub.com.br", "http://www.google.com"),
		html.Title("https://www.youtube.com", "http://www.cod3r.com.br"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
