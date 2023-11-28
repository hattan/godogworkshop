package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Decleration(t *testing.T) {
	// Arrange
	var s string

	// Assert
	assert.Equal(t, "", s)
	assert.NotNil(t, s)
}

func Test_DeclerationAndAssignment(t *testing.T) {
	// Arrange
	var s string
	s = "foo"

	// Assert
	assert.Equal(t, "foo", s)
}

func Test_DeclerationShorthand(t *testing.T) {
	// Arrange
	var s string = "foo"

	// Assert
	assert.Equal(t, "foo", s)
}

func Test_TypeInference(t *testing.T) {
	// Arrange
	// Note shorthand := only works in a function!
	s := "foo"
	i := 1234

	// Assert
	assert.Equal(t, "string", reflect.TypeOf(s).String())
	assert.Equal(t, "int", reflect.TypeOf(i).String())
}

func Test_VariableShadowing(t *testing.T) {
	// Arrange
	var url string
	some_condition := true

	if some_condition {
		url := "https://www.microsoft.com"
		fmt.Println(url)
	} else {
		url := "https://www.google.com"
		fmt.Println(url)
	}

	assert.Equal(t, "", url)
}
