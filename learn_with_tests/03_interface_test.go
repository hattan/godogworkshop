package test

import (
	"testing"

	"github.com/hattan/learngo/pkg/person"
	"github.com/stretchr/testify/assert"
)

func Test_Interface(t *testing.T) {
	// Arrange
	var i interface{}

	// Assert
	assert.Nil(t, i)
}

func Test_Interface_AssertWithString(t *testing.T) {
	// Arrange
	output := runFunctionCaptureStdOut(func() {

		writeLine("test")

	})

	assert.Equal(t, "test\n", output)
}

func Test_Interface_AssertWithInt(t *testing.T) {
	// Arrange
	output := runFunctionCaptureStdOut(func() {
		writeLine(2)
	})

	assert.Equal(t, "2\n", output)
}

func Test_Interface_AssertWithStruct(t *testing.T) {
	// Arrange
	type test struct {
		foo string
	}

	e := test{
		foo: "bar",
	}

	output := runFunctionCaptureStdOut(func() {
		writeLine(e)
	})

	assert.Equal(t, "{bar}\n", output)
}

func Test_Interface_Generics(t *testing.T) {
	// Assert
	assert.Equal(t, 3, sumGeneric(1, 2))
	assert.Equal(t, 5.3, sumGeneric(1.5, 3.8))
}

func Test_Interface_DuckTyping(t *testing.T) {
	// Arrange
	type Named interface {
		FullName() string
	}

	person := person.NewPerson("Bill", "Riker")
	greet := func(n Named) string {
		return "Hello " + n.FullName()
	}

	// Act
	message := greet(person)

	// Assert
	assert.Equal(t, "Hello Bill Riker", message)
}
