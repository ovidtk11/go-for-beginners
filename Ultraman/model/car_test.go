package model

import (
	_ "bytes"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	_ "io"
	_ "log"
	_ "os"
	"testing"
	_ "testing"
)

func Test_Car(t *testing.T) {
	t.Run("ทดสอบ GetEngine", TestCar_GetEngine)
	t.Run("ทดสอบ SetEngine", TestCar_SetEngine)
	t.Run("ทดสอบ GetWheels", TestCar_GetWheels)
	t.Run("ทดสอบ SetWheels", TestCar_SetWheels)
	t.Run("ทดสอบ GetDoors", TestCar_GetDoors)
	t.Run("ทดสอบ SetDoors", TestCar_SetDoors)
	t.Run("ทดสอบ GetSeats", TestCar_GetSeats)
	t.Run("ทดสอบ SetSeats", TestCar_SetSeats)
	t.Run("ทดสอบ GetTopSpeed", TestCar_GetTopSpeed)
	t.Run("ทดสอบ SetTopSpeed", TestCar_SetTopSpeed)

}

func TestCar_GetEngine(t *testing.T) {
	c := car{Engine: "Car"}
	expected := "Car"
	actual := c.GetEngine()
	assert.Equal(t, expected, actual, "is very bad")
}

func TestCar_SetEngine(t *testing.T) {
	c := car{}
	c.SetEngine("Car")
	expected := "Car"
	actual := c.GetEngine()
	assert.Equal(t, expected, actual, "is very bad")
}

func TestCar_GetWheels(t *testing.T) {
	c := car{Wheels: 4}
	expected := 4
	actual := c.GetWheels()
	assert.Equal(t, expected, actual, "is very bad")
}

func TestCar_SetWheels(t *testing.T) {
	c := car{}
	c.SetWheels(4)
	expected := 4
	actual := c.GetWheels()
	assert.Equal(t, expected, actual, "is very bad")
}

func TestCar_GetDoors(t *testing.T) {
	c := car{Doors: 4}
	expected := 4
	actual := c.GetDoors()
	assert.Equal(t, expected, actual, "is very bad")
}

func TestCar_SetDoors(t *testing.T) {
	c := car{}
	c.SetDoors(4)
	expected := 4
	actual := c.GetDoors()
	assert.Equal(t, expected, actual, "is very bad")
}

func TestCar_GetSeats(t *testing.T) {
	c := car{Seats: 4}
	expected := 4
	actual := c.GetSeats()
	assert.Equal(t, expected, actual, "is very bad")
}

func TestCar_SetSeats(t *testing.T) {
	c := car{}
	c.SetSeats(4)
	expected := 4
	actual := c.GetSeats()
	assert.Equal(t, expected, actual, "is very bad")
}

func TestCar_GetTopSpeed(t *testing.T) {
	c := car{TopSpeed: 4}
	expected := 4
	actual := c.GetTopSpeed()
	assert.Equal(t, expected, actual, "is very bad")
}

func TestCar_SetTopSpeed(t *testing.T) {
	c := car{}
	c.SetTopSpeed(4)
	expected := 4
	actual := c.GetTopSpeed()
	assert.Equal(t, expected, actual, "is very bad")
}
