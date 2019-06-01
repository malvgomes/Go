package main

import "fmt"

type sport interface {
	startTurbo()
}

type ferrari struct {
	model   string
	turboOn bool
	curVel  int
}

func (f *ferrari) startTurbo() {
	f.turboOn = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.startTurbo()

	var car2 sport = &ferrari{"F40", false, 0}
	car2.startTurbo()

	fmt.Println(car1, car2)
}
