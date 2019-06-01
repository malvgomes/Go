package main

import (
	"fmt"
	"time"
)

func speak(person, text string, qtt int) {
	for i := 0; i < qtt; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i+1)
	}
}

func main() {
	// speak("Maria", "Why don't you speak to me?", 3)
	// speak("John", "I can only speak after you", 1)

	// go speak("Maria", "Hey", 500)
	// go speak("John", "Hello", 500)
	// time.Sleep(time.Second * 5)

	go speak("Maria", "Understood", 10)
	speak("John", "Congrats", 5)

	fmt.Println("End")
}
