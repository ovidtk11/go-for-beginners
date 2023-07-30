package model

import "fmt"

type animal struct {
	Name   string
	Age    int
	Height float64
	Sound  string
}

func (a *animal) Greeting() {
	fmt.Printf("Hello, I‘m %s \n", a.Name)
}

func (a *animal) GetSound() string {
	return a.Sound
}

func (a *animal) SetSound(sound string) {
	a.Sound = sound
}

func (a *animal) MakeSound() {
	fmt.Printf("%s says %s\n", a.Name, a.Sound)
}

func (a *animal) GetAge() int {
	return a.Age
}

func (a *animal) SetAge(i int) {
	a.Age = i
}

func (a *animal) GetHeight() float64 {
	return a.Height
}

func (a *animal) SetHeight(f float64) {
	a.Height = f
}

func (a *animal) GetName() string {
	return a.Name
}

func (a *animal) SetName(s string) {
	a.Name = s
}

// เราทำ interface เพื่อเป็นตัวกลางในการเชื่อมต่อไปยัง object,function,variable ที่เป็น private และอยู่คนละ package ได้

type IAnimal interface {
	GetName() string
	SetName(s string)
	GetAge() int
	SetAge(i int)
	GetHeight() float64
	SetHeight(f float64)
	MakeSound()
	GetSound() string
	SetSound(sound string)
	Greeting()
}

// NewAnimal เราสร้าง function นี้ขึ้นมาเพื่อทำให้  animal{}
// implement ความสามารถตาม interface IAnimal
// เพื่อให้ package อื่นที่ทำการเรียก NewAnimal ได้รับค่าของ animal{}
// ได้รับค่าของ struct animal{} และความสามารถ IAnimal
// ในการ interact กับ private animal{} นั่นเอง
func NewAnimal() IAnimal {
	return &animal{}
}

//1 สร้าง struct
//2 สร้าง interface
//3 ทำให้ struct มีความสามารถตาม interface ที่เราสร้างไว้ (implement)
