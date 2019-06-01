package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	fstName string
	lstName string
}

type product struct {
	name  string
	price float64
}

//interfaces are implicitly implemented
func (p person) toString() string {
	return p.fstName + " " + p.lstName
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var thing printable = person{"Matheus", "Alves"}
	fmt.Println(thing.toString())
	print(thing)

	thing = product{"Lápis", 2.00}
	print(thing)
}
