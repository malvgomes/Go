package main

import "fmt"

type Fahrenheit float64
type Celsius float64

//Type method
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {
	var c Celsius = -273.0

	fmt.Println(c.String())
}
