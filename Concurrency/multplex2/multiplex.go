package main

import (
	"fmt"
	"time"
)

func speak(person string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s speaking %d", person, i)
		}
	}()

	return c
}

func join(en1, en2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-en1:
				c <- s
			case s := <-en2:
				c <- s
			}
		}
	}()

	return c
}

func main() {
	c := join(speak("John"), speak("Mary"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
