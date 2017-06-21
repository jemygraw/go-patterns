package main

import (
	"fmt"
)

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Interface interface {
	Drive() error
	Stop() error
}

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Car struct {
	color    Color
	wheels   Wheels
	topSpeed Speed
}

func NewCarBuilder() Builder {
	return &Car{}
}

func (c *Car) Color(color Color) Builder {
	c.color = color
	return c
}

func (c *Car) Wheels(wheels Wheels) Builder {
	c.wheels = wheels
	return c
}

func (c *Car) TopSpeed(topSpeed Speed) Builder {
	c.topSpeed = topSpeed
	return c
}
func (c *Car) Build() Interface {
	return c
}

func (c *Car) Drive() error {
	fmt.Println("Drive the car")
	fmt.Println("Color:", c.color)
	fmt.Println("Wheels:", c.wheels)
	fmt.Println("TopSpeed:", c.topSpeed)
	fmt.Println("----------")
	return nil
}

func (c *Car) Stop() error {
	fmt.Println("Stop the car")
	return nil
}

func main() {
	assembly := NewCarBuilder().Color(RedColor)

	familyCar := assembly.Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	sportsCar.Drive()
}
