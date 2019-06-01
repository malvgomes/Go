package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct to json
	p := product{1, "Notebook", 2600.00, []string{"Deals", "Eletronics"}}
	pToJSON, _ := json.Marshal(p)
	fmt.Println(string(pToJSON))

	var p2 product
	js := `
		{
			"id":1,
			"name":"Notebook",
			"price":2600,
			"tags":["Deals","Eletronics"]
		}`
	json.Unmarshal([]byte(js), &p2)

	fmt.Println(p2)
}
