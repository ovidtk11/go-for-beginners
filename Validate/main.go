package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// fmt.Println("Hello, Ovi")

	// s := Student{
	// 	FirstName: "Ovi",
	// 	LastName:  "Tongkul",
	// 	Age: 12,
	// }

	// b, _ := json.Marshal(s)
	// fmt.Println(string(b))

	// s.Validate()

	// =================================

	car := Car{
		Brand:  "Toyota",
		Model:  "Corolla",
		Door:   4,
		Tire:   4,
		Wheels: 4,
	}

	b, _ := json.Marshal(car)
	fmt.Println(string(b))

	car.Validate()
}

type Student struct {
	FirstName string
	LastName  string
	Age       int
}

func (s *Student) Validate() {
	s.ValidateFirstName()
	s.ValidateLastName()
	s.ValidateAge()
}

func (s *Student) ValidateFirstName() {
	// create regex
	regex := regexp.MustCompile(`^[ก-๏]+$`)

	// check
	if !regex.MatchString(s.FirstName) {
		fmt.Println("FirstName ไม่ถูกต้อง")
		os.Exit(0)
	}

}

func (s *Student) ValidateLastName() {
	// create regex
	regex := regexp.MustCompile(`^[ก-๏]+$`)

	// check
	if !regex.MatchString(s.LastName) {
		fmt.Println("LastName ไม่ถูกต้อง")
		os.Exit(0)
	}
}

func (s *Student) ValidateAge() {
	// create regex
	regex := regexp.MustCompile(`^(1\d|20|[1-9])$|^[2-9]\d$`)

	// check
	if !regex.MatchString(strconv.Itoa(s.Age)) {
		fmt.Println("Age ไม่ถูกต้อง")
		os.Exit(0)
	}

}

type Car struct {
	Brand  string
	Model  string
	Door   int
	Tire   int
	Wheels int
}

func (c *Car) Validate() string {

	err := ""

	err = c.validateBrand()
	if err != "" {
		return err
	}
	
	err = c.validateModel()
	if err != "" {
		return err
	}

	err = c.validateDoor()
	if err != "" {
		return err
	}

	err = c.validateTire()
	if err != "" {
		return err
	}

	err = c.validateWheels()
	if err != "" {
		return err
	}

	return "No Error"
}

func (c *Car) validateBrand() string {
	if c.Brand == "" {
		return "Brand ไม่ถูกต้อง"
	}
	return ""
}

func (c *Car) validateModel() string {
	if c.Model == "" {
		return "Model ไม่ถูกต้อง"
	}
	return ""
}

func (c *Car) validateDoor() string {
	if c.Door == 0 {
		return "Door ไม่ถูกต้อง"
	}
	return ""
}

func (c *Car) validateTire() string {
	if c.Tire == 0 {
		return "Tire ไม่ถูกต้อง"
	}
	return ""
}

func (c *Car) validateWheels() string {
	if c.Wheels == 0 {
		return "Wheels ไม่ถูกต้อง"
	}
	return ""
}



func ABC() string {
	
	a := "A"
	b := "B"
	c := "C"

	ABC := a + b + c

	return ABC
}

func DEF() (message string) {
	name := "ovi"
	msg := ""

	if name == "Ovi" {
		msg = "Hello, Ovi"
	}

	return msg
}

func HIJ() (message string) {
	name := "ovi"

	if name == "Ovi" {
		return "Hello, Ovi"
	}

	return ""
}