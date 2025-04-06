package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck interface {
	loadCargo() error
	unloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

// A method is a function with a receiver arguement
func (t *NormalTruck) loadCargo() error {
	t.cargo += 1
	return nil
}

func (t *NormalTruck) unloadCargo() error {
	t.cargo = 0
	return nil
}

// processTruck() handles the loading and unloading of a truck
func processTruck(truck Truck) error {
	fmt.Printf("started processing truck: %+v\n", truck)

	//Processing time
	time.Sleep(time.Second)

	if err := truck.loadCargo(); err != nil {
		return fmt.Errorf("cargo not loaded: %w", err)
	}
	if err := truck.unloadCargo(); err != nil {
		return fmt.Errorf("cargo not loaded: %w", err)
	}

	fmt.Printf("finished processing truck: %+v\n", truck)

	return nil
}

// Implemented loadCargo for type ElectricTruck
func (e *ElectricTruck) loadCargo() error {
	e.cargo += 1
	e.battery -= 1
	return nil
}

func (e *ElectricTruck) unloadCargo() error {
	e.cargo = 0
	e.battery -= 1
	return nil
}

func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup
	for _, t := range trucks {

		wg.Add(1)

		go func(t Truck) {
			processTruck(t)
			wg.Done()
		}(t)

	}
	wg.Wait()

	return nil
}

func main() {
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Printf("All trucks processed successfully!")
}
