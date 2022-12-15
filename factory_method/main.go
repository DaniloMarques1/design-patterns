package main

import (
	"fmt"
	"log"
)

type Shipping interface {
	Deliver() error
}

type Truck struct {
}

func (t Truck) Deliver() error {
	fmt.Println("Delivering by truck")
	return nil
}

type Ship struct {
}

func (s Ship) Deliver() error {
	fmt.Println("Delivering by ship")
	return nil
}

type ShippingDelivery interface {
	CreateShipping() Shipping
	Deliver() error
}

type ShippingByTruck struct {
}

func (s ShippingByTruck) CreateShipping() Shipping {
	return Truck{}
}

func (s ShippingByTruck) Deliver() error {
	shipping := s.CreateShipping()
	shipping.Deliver()
	return nil
}

type ShippingByShip struct {
}

func (s ShippingByShip) CreateShipping() Shipping {
	return Ship{}
}

func (s ShippingByShip) Deliver() error {
	shipping := s.CreateShipping()
	shipping.Deliver()
	return nil
}

func main() {
	var shipping ShippingDelivery
	shippingType := "ship"
	switch shippingType {
	case "truck":
		shipping = ShippingByTruck{}
	case "ship":
		shipping = ShippingByShip{}
	default:
		log.Fatalf("Wrong shipping type\n")
	}

	shipping.Deliver()
	
}
