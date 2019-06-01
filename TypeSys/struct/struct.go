package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

//Method: func with receiver
func (p product) getName() string {
	return p.name
}

func (p product) getPrice() float64 {
	return p.price
}

func (p product) getDiscount() float64 {
	return p.discount
}

func (p product) getPriceWDiscount() float64 {
	return (1 - p.discount) * p.price
}

func main() {
	pencil := product{
		name:     "Pencil",
		price:    2.50,
		discount: 0.05,
	}

	fmt.Println(pencil.getName())
	fmt.Println(pencil.getPrice())
	fmt.Println(pencil.getDiscount())
	fmt.Println(pencil.getPriceWDiscount())
}
