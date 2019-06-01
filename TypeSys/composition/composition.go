package main

import "fmt"

type sport interface {
	startTurbo()
}

type luxury interface {
	autoParking()
}

type sportLux interface {
	sport
	luxury
	//and more
}

type bmw7 struct{}

func (b bmw7) startTurbo() {
	fmt.Println("turbo")
}

func (b bmw7) autoParking() {
	fmt.Println("parking")
}

func main() {
	var b sportLux = bmw7{}
	b.autoParking()
	b.startTurbo()
}
