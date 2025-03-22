package main

import "fmt"

type Truck struct {
	id string
}

func processTruck(truck Truck) {
	fmt.Printf("Processing truck : %s\n", truck.id)
}

func main() {
	Trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
		{id: "Truck-4"},
	}

	for _, truck := range Trucks {
		fmt.Printf("Truck %s arrived\n", truck.id)
		processTruck(truck)
	}

}
