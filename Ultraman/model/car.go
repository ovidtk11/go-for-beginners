package model

import _ "fmt"

type car struct {
	Engine   string
	Wheels   int
	Doors    int
	Seats    int
	TopSpeed int
}

func (c *car) SetEngine(engine string) {
	c.Engine = engine
}

func (c *car) SetWheels(wheels int) {
	c.Wheels = wheels
}

func (c *car) SetDoors(doors int) {
	c.Doors = doors
}

func (c *car) SetSeats(seats int) {
	c.Seats = seats
}

func (c *car) SetTopSpeed(topSpeed int) {
	c.TopSpeed = topSpeed
}

func (c *car) GetEngine() string {
	return c.Engine
}

func (c *car) GetWheels() int {
	return c.Wheels
}

func (c *car) GetDoors() int {
	return c.Doors
}

func (c *car) GetSeats() int {
	return c.Seats
}

func (c *car) GetTopSpeed() int {
	return c.TopSpeed
}

type ICar interface {
	ICarSetter
	ICarGetter
}
type ICarGetter interface {
	GetEngine() string
	GetWheels() int
	GetDoors() int
	GetSeats() int
	GetTopSpeed() int
}
type ICarSetter interface {
	SetEngine(engine string)
	SetWheels(wheels int)
	SetDoors(doors int)
	SetSeats(seats int)
	SetTopSpeed(topSpeed int)
}

func NewCar() ICar {
	return &car{}
}
