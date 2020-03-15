package main

import (
	"fmt"
)

type Car struct {
	Name      string
	Model     string
	TopSpeed  int
	TyresType string
}

func (c Car) GiveInfo() {
	fmt.Printf("Car info:\n Name: %s\n Model: %s\n TopSpeed: %b\n TyresType: %s\n", c.Name, c.Model, c.TopSpeed, c.TyresType)
}

func CarName(c *Car) {
	fmt.Printf("Car name is %s\n", c.Name)
}

func NewCar(name string, model string, speed int, tyres string) *Car {
	return &Car{
		Name:      name,
		Model:     model,
		TopSpeed:  speed,
		TyresType: tyres,
	}
}

func carDecorator(f func(c *Car)) func(*Car) {
	return func(c *Car) {
		fmt.Printf("Model => %s\n", c.Model)
		if c.Model == "Toyota" {
			fmt.Println("The price is 1000000$")
		} else if c.Model == "Mazda" {
			fmt.Println("The price is 200000$")
		}
		f(c)
	}
}

func main() {
	car1 := NewCar("Booboo", "Toyota", 300, "Soft")
	car2 := NewCar("Beebee", "Mazda", 200, "Hard")
	car3 := NewCar("Buubuu", "Ferrari", 150, "Medium")

	car1.GiveInfo()
	car2.GiveInfo()
	CarName(car3)

	car4 := NewCar("Booboo", "Toyota", 300, "Soft")

	fmt.Printf("=============DECORATOR=================\n")
	carDecorator(CarName)(car4)
	carDecorator(CarName)(car2)

}
