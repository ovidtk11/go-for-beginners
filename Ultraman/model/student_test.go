package model

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	_ "log"
	"os"
	"testing"
	_ "testing"
)

func Test_Student(t *testing.T) {
	t.Run("ทดสอบ GetFirstName", test_GetFirstName)
	t.Run("ทดสอบ SetFirstName", test_SetFirstName)
	t.Run("ทดสอบ GetLastName", test_GetLastName)
	t.Run("ทดสอบ SetLastName", test_SetLastName)
	t.Run("ทดสอบ GetAge", test_GetAge)
	t.Run("ทดสอบ SetAge", test_SetAge)
	t.Run("ทดสอบ GetHeight", test_GetHeight)
	t.Run("ทดสอบ SetHeight", test_SetHeight)
	t.Run("ทดสอบ GetWeight", test_GetWeight)
	t.Run("ทดสอบ SetWeight", test_SetWeight)
	t.Run("ทดสอบ Greeting", test_Greeting)

}

func test_GetFirstName(t *testing.T) {
	s := student{FirstName: "ovi"}
	expected := "ovi"
	actual := s.GetFirstName()
	assert.Equal(t, expected, actual, "is bad")
}

func test_SetFirstName(t *testing.T) {
	s := student{}
	s.SetFirstName("ovi")
	expected := "ovi"
	actual := s.GetFirstName()
	assert.Equal(t, expected, actual, "is bad")
}

func test_GetLastName(t *testing.T) {
	s := student{LastName: "ovi"}
	expected := "ovi"
	actual := s.GetLastName()
	assert.Equal(t, expected, actual, "is bad")
}

func test_SetLastName(t *testing.T) {
	s := student{}
	s.SetLastName("ovi")
	expected := "ovi"
	actual := s.GetLastName()
	assert.Equal(t, expected, actual, "is bad")
}

func test_GetAge(t *testing.T) {
	s := student{Age: 12}
	expected := 12
	actual := s.GetAge()
	assert.Equal(t, expected, actual, "is bad")
}

func test_SetAge(t *testing.T) {
	s := student{}
	s.SetAge(12)
	expected := 12
	actual := s.GetAge()
	assert.Equal(t, expected, actual, "is bad")
}
func test_GetHeight(t *testing.T) {
	s := student{Height: 160}
	expected := 160.00
	actual := s.GetHeight()
	assert.Equal(t, expected, actual, "is bad")
}

func test_SetHeight(t *testing.T) {
	s := student{}
	s.SetHeight(160)
	expected := 160.00
	actual := s.GetHeight()
	assert.Equal(t, expected, actual, "is bad")
}

func test_GetWeight(t *testing.T) {
	s := student{Weight: 38}
	expected := 38.00
	actual := s.GetWeight()
	assert.Equal(t, expected, actual, "is bad")
}

func test_SetWeight(t *testing.T) {
	s := student{}
	s.SetWeight(38)
	expected := 38.00
	actual := s.GetWeight()
	assert.Equal(t, expected, actual, "is bad")
}

func test_Greeting(t *testing.T) {

	s := student{FirstName: "ovi"}

	actual := CaptureStdout(s.Greeting)

	expected := "Hello, I‘m ovi \n"

	assert.Equal(t, expected, actual)
}

func CaptureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	r.Close()

	return buf.String()
}
