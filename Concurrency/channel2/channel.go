package main

import (
	"fmt"
	"time"
)

func fourTimes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go fourTimes(2, c)

	a, b := <-c, <-c // receiving channel data
	fmt.Println(a, b)

	fmt.Println(<-c)
}
