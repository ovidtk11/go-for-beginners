package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Car(t *testing.T) {
	t.Run("Happy-Case: ทดสอบการ Validate Brand", validateBrand)
	t.Run("Sad-Case: ทดสอบการ Validate Brand", validateBrandInvalid)

	t.Run("Happy-Case: ทดสอบการ Validate Model", validateModel)
	t.Run("Sad-Case: ทดสอบการ Validate Model", validateModelInvalid)

	t.Run("Happy-Case: ทดสอบการ Validate Door", validateDoor)
	t.Run("Sad-Case: ทดสอบการ Validate Door", validateDoorInvalid)

	t.Run("Happy-Case: ทดสอบการ Validate Tire", validateTire)
	t.Run("Sad-Case: ทดสอบการ Validate Tire", validateTireInvalid)

	t.Run("Happy-Case: ทดสอบการ Validate Wheels", validateWheels)
	t.Run("Sad-Case: ทดสอบการ Validate Wheels", validateWheelsInvalid)
}

func validateBrand(t *testing.T) {
	car := Car{
		Brand:  "Toyota",
		Model:  "Corolla",
		Door:   4,
		Tire:   4,
		Wheels: 4,
	}

	expected := "No Error"
	actual := car.Validate()

	assert.Equal(t, expected, actual)
}


func validateModel(t *testing.T) {
	car := Car{
		Brand:  "Toyota",
		Model:  "Corolla",
		Door:   4,
		Tire:   4,
		Wheels: 4,
	}

	expected := "No Error"
	actual := car.Validate()

	assert.Equal(t, expected, actual)
}

func validateDoor(t *testing.T) {
	car := Car{
		Brand:  "Toyota",
		Model:  "Corolla",
		Door:   4,
		Tire:   4,
		Wheels: 4,
	}

	expected := "No Error"
	actual := car.Validate()

	assert.Equal(t, expected, actual)
}

func validateTire(t *testing.T) {
	car := Car{
		Brand:  "Toyota",
		Model:  "Corolla",
		Door:   4,
		Tire:   4,
		Wheels: 4,
	}

	expected := "No Error"
	actual := car.Validate()

	assert.Equal(t, expected, actual)
}

func validateWheels(t *testing.T) {
	car := Car{
		Brand:  "Toyota",
		Model:  "Corolla",
		Door:   4,
		Tire:   4,
		Wheels: 4,
	}

	expected := "No Error"
	actual := car.Validate()

	assert.Equal(t, expected, actual)
}

func validateBrandInvalid(t *testing.T) {
	car := Car{
		Brand:  "",
		Model:  "Corolla",
		Door:   4,
		Tire:   4,
		Wheels: 4,
	}

	expected := "Brand ไม่ถูกต้อง"
	actual := car.Validate()

	assert.Equal(t, expected, actual)
}

func validateModelInvalid(t *testing.T) {
	car := Car{
		Brand:  "Toyota",
		Model:  "",
		Door:   4,
		Tire:   4,
		Wheels: 4,
	}

	expected := "Model ไม่ถูกต้อง"
	actual := car.Validate()

	assert.Equal(t, expected, actual)
}

func validateDoorInvalid(t *testing.T) {
	car := Car{
		Brand:  "Toyota",
		Model:  "Corolla",
		Door:   0,
		Tire:   4,
		Wheels: 4,
	}

	expected := "Door ไม่ถูกต้อง"
	actual := car.Validate()

	assert.Equal(t, expected, actual)
}

func validateTireInvalid(t *testing.T) {
	car := Car{
		Brand:  "Toyota",
		Model:  "Corolla",
		Door:   4,
		Tire:   0,
		Wheels: 4,
	}

	expected := "Tire ไม่ถูกต้อง"
	actual := car.Validate()

	assert.Equal(t, expected, actual)
}

func validateWheelsInvalid(t *testing.T) {
	car := Car{
		Brand:  "Toyota",
		Model:  "Corolla",
		Door:   4,
		Tire:   4,
		Wheels: 0,
	}

	expected := "Wheels ไม่ถูกต้อง"
	actual := car.Validate()

	assert.Equal(t, expected, actual)
}

