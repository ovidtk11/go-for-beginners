package main

import (
	"fmt"
	"struct/student"
)

func main() {
	s1 := student.NewStudent()
	s1.SetFirstName("Ovi")

	s2 := student.NewStudent()
	s2.SetFirstName("Noah")

	fmt.Printf("s1 name: => %v\n", s1.GetFirstName())

	fmt.Printf("s2 name: => %v\n", s2.GetFirstName())

}
