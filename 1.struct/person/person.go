package person

import "fmt"

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p Person) Introduce() string {
	return fmt.Sprintf("Hello, I'm %s and I'm %d years old.\n", p.Name, p.Age)
}
