package main

import (
	"fmt"
	"struct/student"
)

func main() {
	s := student.Student{}

	fmt.Printf("Print without fields ==> %v\n", s)
	// print struct with fields
	fmt.Printf("Print with fields ==> %#v\n", s)

}
