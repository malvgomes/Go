package main

import "fmt"

type item struct {
	productID int
	quantity  int
	price     float64
}

type order struct {
	userID int
	items  []item
}

func (o order) orderTotal() float64 {
	total := 0.0
	for _, ord := range o.items {
		total += ord.price
	}

	return total
}

func main() {
	o := order{
		userID: 1,
		items: []item{
			item{productID: 1, quantity: 10, price: 10.0},
			item{productID: 2, quantity: 25, price: 13.0},
			item{productID: 3, quantity: 22, price: 9.25},
		},
	}

	fmt.Println(o)
	fmt.Println(o.orderTotal())
}
