package model

import _ "fmt"

type godzilla struct {
	Name     string  //*Godzilla Earth
	Age      int     //20000
	Height   float64 //300
	Weight   float64 //100000
	Ultimate string  //Atomic breath
}

func (g *godzilla) GetUltimate() string {
	return g.Ultimate
}

func (g *godzilla) SetUltimate(Ultimate string) {
	g.Ultimate = Ultimate
}

func (g *godzilla) GetName() string {
	return g.Name
}

func (g *godzilla) SetName(Name string) {
	g.Name = Name
}

func (g *godzilla) GetAge() int {
	return g.Age
}

func (g *godzilla) SetAge(Age int) {
	g.Age = Age
}

func (g *godzilla) GetHeight() float64 {
	return g.Height
}

func (g *godzilla) SetHeight(Height float64) {
	g.Height = Height
}

func (g *godzilla) GetWeight() float64 {
	return g.Weight
}

func (g *godzilla) SetWeight(Weight float64) {
	g.Weight = Weight
}

type IGodzilla interface {
	GetName() string
	SetName(Name string)
	GetAge() int
	SetAge(Age int)
	GetHeight() float64
	SetHeight(Height float64)
	GetWeight() float64
	SetWeight(Weight float64)
	GetUltimate() string
	SetUltimate(Ultimate string)
}

func NewGodzilla() IGodzilla {
	return &godzilla{}
}
