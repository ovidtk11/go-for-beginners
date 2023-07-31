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

func Test_Godzilla(t *testing.T) {
	t.Run("ทดสอบ GetName", test_GetName_godzilla)
	t.Run("ทดสอบ SetName", test_SetName_godzilla)
	t.Run("ทดสอบ GetAge_godzilla", test_GetAge_godzilla)
	t.Run("ทดสอบ SetAge_godzilla", test_SetAge_godzilla)
	t.Run("ทดสอบ GetHeight_godzilla", test_GetHeight_godzilla)
	t.Run("ทดสอบ SetHeight_godzilla", test_SetHeight_godzilla)
	t.Run("ทดสอบ SetWeight_godzilla", test_GetWeight_godzilla)
	t.Run("ทดสอบ SetWeight_godzilla", test_SetWeight_godzilla)
	t.Run("ทดสอบ SetUltimate_godzilla", test_GetUltimate_godzilla)
	t.Run("ทดสอบ SetUltimate  _godzilla", test_SetUltimate_godzilla)
}

func test_GetName_godzilla(t *testing.T) {
	g := godzilla{Name: "Godzilla Earth"}
	expected := "Godzilla Earth"
	actual := g.GetName()
	assert.Equal(t, expected, actual, "is very bad")
}

func test_SetName_godzilla(t *testing.T) {
	g := godzilla{}
	g.SetName("Godzilla Earth")
	expected := "Godzilla Earth"
	actual := g.GetName()
	assert.Equal(t, expected, actual, "is very bad")
}

func test_GetAge_godzilla(t *testing.T) {
	g := godzilla{Age: 20000}
	expected := 20000
	actual := g.GetAge()
	assert.Equal(t, expected, actual, "is very bad")
}

func test_SetAge_godzilla(t *testing.T) {
	g := godzilla{}
	g.SetAge(20000)
	expected := 20000
	actual := g.GetAge()
	assert.Equal(t, expected, actual, "is very bad")
}

func test_GetHeight_godzilla(t *testing.T) {
	g := godzilla{Height: 300}
	expected := 300.00
	actual := g.GetHeight()
	assert.Equal(t, expected, actual, "is very bad")
}

func test_SetHeight_godzilla(t *testing.T) {
	g := godzilla{}
	g.SetHeight(300)
	expected := 300.00
	actual := g.GetHeight()
	assert.Equal(t, expected, actual, "is very bad")
}

func test_GetWeight_godzilla(t *testing.T) {
	g := godzilla{Weight: 100000}
	expected := 100000.00
	actual := g.GetWeight()
	assert.Equal(t, expected, actual, "is very bad")
}

func test_SetWeight_godzilla(t *testing.T) {
	g := godzilla{}
	g.SetWeight(100000)
	expected := 100000.00
	actual := g.GetWeight()
	assert.Equal(t, expected, actual, "is very bad")
}

func test_GetUltimate_godzilla(t *testing.T) {
	g := godzilla{Ultimate: "Atomic breath"}
	expected := "Atomic breath"
	actual := g.GetUltimate()
	assert.Equal(t, expected, actual, "is very bad")
}

func test_SetUltimate_godzilla(t *testing.T) {
	g := godzilla{}
	g.SetUltimate("Atomic breath")
	expected := "Atomic breath"
	actual := g.GetUltimate()
	assert.Equal(t, expected, actual, "is very bad")
}
