package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //lock operation
	fmt.Println("Only after read")
}

func main() {
	c := make(chan int) //no buffer

	go routine(c)
	fmt.Println(<-c) //lock operation
	fmt.Println("Was read")

	fmt.Println(<-c) //deadlock
	fmt.Println("Fim")
}
