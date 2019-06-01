package main

import "fmt"

type car struct {
	name   string
	curVel int
}

type ferrari struct {
	car     //anonym fields
	turboOn bool
}

func (c car) getName() string {
	return c.name
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.curVel = 0
	f.turboOn = true

	fmt.Printf("Ferrari %s has its turbo on? %v\n", f.getName(), f.turboOn)
}
