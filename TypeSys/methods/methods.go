package main

import (
	"fmt"
	"strings"
)

type person struct {
	fstName string
	lstName string
}

func (p person) getFullName() string {
	return p.fstName + " " + p.lstName
}

func (p *person) setFullName(fullName string) {
	names := strings.Split(fullName, " ")
	p.fstName = names[0]
	p.lstName = names[1]
}

func (p person) greetings() {
	fmt.Println("Hello! My name is " + p.getFullName())
}
func main() {
	p1 := person{"Peter", "Jackson"}
	p1.setFullName("Matheus Alves")
	p1.greetings()
}
