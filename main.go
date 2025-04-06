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
	fmt.Printf("processing truck: %+v\n", truck)
	if err := truck.loadCargo(); err != nil {
		return fmt.Errorf("cargo not loaded: %w", err)
	}
	if err := truck.unloadCargo(); err != nil {
		return fmt.Errorf("cargo not loaded: %w", err)
	}
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

func main() {
	nt := &NormalTruck{id: "1"}
	if err := processTruck(nt); err != nil {
		log.Fatalf("error processing truck: %s", err)
	}

	et := &ElectricTruck{id: "1"}
	if err := processTruck(et); err != nil {
		log.Fatalf("error processing truck: %s", err)
	}

	fmt.Println(nt.cargo)
	fmt.Println(et.battery)
}
