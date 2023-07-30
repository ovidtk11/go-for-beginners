package model

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	a := NewAnimal()
	a.SetName("ovi")
	a.SetSound("meow")
	a.MakeSound()
	a.Greeting()
}

var a string

func Test_ABC(t *testing.T) {

	a = "ABC"

	log.Println("Before swap:", a)

	swap(&a)

	log.Println("After swap:", a)

}

func swap(s *string) {
	*s = "CBA"
}
