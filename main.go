package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck struct {
	id string
}

// A method is a function with a receiver arguement
func (t Truck) loadCargo() error {
	return ErrTruckNotFound
}

// processTruck() handles the loading and unloading of a truck
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck : %s\n", truck.id)
	if err := truck.loadCargo(); err != nil {
		return fmt.Errorf("cargo not loaded: %w", err)
	}
	return ErrNotImplemented
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

		if err := processTruck(truck); err != nil {
			log.Fatalf("error processing truck: %s", err)
		}

	}

}
