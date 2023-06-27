package main

import (
	"fmt"

	"github.com/ovidtk11/go-for-beginners/1.variables/person"
)

func main() {

	p := new(person.Person)
	p.Name = "Ovi"
	p.Age = 11

	p2 := new(person.Person)
	p2.Name = "Noah"
	p2.Age = 7

	p3 := new(person.Person)
	p3.Name = "Jayneen"
	p3.Age = 3

	ps := make([]person.Person, 0)

	ps = append(ps, *p, *p2, *p3)

	for _, p := range ps {
		fmt.Println(p.String())
		fmt.Println(p.Introduce())
	}
}
