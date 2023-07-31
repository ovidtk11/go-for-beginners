package model

import (
	"fmt"
	_ "fmt"
)

type student struct {
	FirstName string
	LastName  string
	Age       int
	Height    float64
	Weight    float64
}

func (s *student) Greeting() {
	fmt.Printf("Hello, I‘m %s \n", s.FirstName)
}

func (s *student) GetFirstName() string {
	return s.FirstName
}

func (s *student) SetFirstName(firstName string) {
	s.FirstName = firstName
}

func (s *student) GetLastName() string {
	return s.LastName
}

func (s *student) SetLastName(lastName string) {
	s.LastName = lastName
}

func (s *student) GetAge() int {
	return s.Age
}

func (s *student) SetAge(age int) {
	s.Age = age
}

func (s *student) GetHeight() float64 {
	return s.Height
}

func (s *student) SetHeight(height float64) {
	s.Height = height
}

func (s *student) GetWeight() float64 {
	return s.Weight
}

func (s *student) SetWeight(weight float64) {
	s.Weight = weight
}

type IStudent interface {
	GetFirstName() string
	SetFirstName(firstName string)
	GetLastName() string
	SetLastName(lastName string)
	GetAge() int
	SetAge(age int)
	GetHeight() float64
	SetHeight(height float64)
	GetWeight() float64
	SetWeight(weight float64)
	Greeting()
}

func NewStudent() IStudent {
	return &student{}
}
