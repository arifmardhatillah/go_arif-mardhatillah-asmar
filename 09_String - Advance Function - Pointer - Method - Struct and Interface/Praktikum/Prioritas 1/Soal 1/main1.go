package main

import "fmt"

type Car struct {
	typeCar string
	fuelIn  int
}

func (e *Car) perkiraanJarak() float64 {
	return float64(e.fuelIn) / (1.5)
}

func main() {
	minibus := Car{
		typeCar: "Minibus",
		fuelIn:  15,
	}

	fmt.Println("Jarak yang ditempuh:", minibus.perkiraanJarak(), "mill")
}
