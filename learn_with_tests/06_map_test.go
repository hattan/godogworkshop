package test

import (
	"fmt"
	"maps"
	"testing"

	"github.com/hattan/learngo/pkg/wcase"
	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	// Arrange
	var m map[string]int

	// Assert
	assert.Nil(t, m)
}

func Test_Map_Make(t *testing.T) {
	// Arrange
	var m map[string]int

	// Act
	m = make(map[string]int)
	m["foo"] = 12345

	// Assert
	assert.NotNil(t, m)
	assert.Equal(t, 1, len(m))
	assert.Equal(t, 12345, m["foo"])
}

func Test_Map_Literal(t *testing.T) {
	// Arrange
	m := map[string]int{
		"foo": 12345,
		"bar": 67890,
	}

	// Assert
	assert.NotNil(t, m)
	assert.Equal(t, 2, len(m))
	assert.Equal(t, 12345, m["foo"])
	assert.Equal(t, 67890, m["bar"])
}

func Test_Map_Delete(t *testing.T) {
	// Arrange
	m := map[string]int{
		"foo": 12345,
		"bar": 67890,
	}

	// Act
	delete(m, "foo")

	// Assert
	assert.NotNil(t, m)
	assert.Equal(t, 1, len(m))
	assert.NotEqual(t, 12345, m["foo"])
	assert.Equal(t, 67890, m["bar"])
}

func Test_Map_CheckElementExists(t *testing.T) {
	// Arrange
	m := map[string]int{
		"foo": 12345,
		"bar": 67890,
	}

	// Act
	delete(m, "foo")
	v, ok := m["foo"]

	// Assert
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func Test_Map_MapsDelete(t *testing.T) {
	// Arrange
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// Act
	maps.DeleteFunc(m, func(k string, v int) bool {
		return v%2 == 0 // delete odd values
	})

	// Assert
	v, ok := m["b"]
	assert.False(t, ok)
	assert.Equal(t, 0, v)

}

func Test_Map_WordCountExcercise(t *testing.T) {
	// Arrange
	input := "I am learning I Go!"

	// Act
	m := WordCount(input, wcase.SENSITIVE)

	// Assert
	expected := map[string]int{
		"I":        2,
		"am":       1,
		"learning": 1,
		"Go!":      1,
	}

	assert.NotNil(t, m)
	for i, v := range expected {
		assert.Equal(t, v, m[i])
	}

}

func Test_Map_WordCountExcerciseTableDrive(t *testing.T) {
	type test struct {
		input    string
		expected map[string]int
	}

	tests := []test{
		{input: "I am learning go!", expected: map[string]int{"i": 1, "am": 1, "learning": 1, "go!": 1}},
		{input: "The quick brown fox jumped over the lazy dog", expected: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumped": 1, "over": 1, "lazy": 1, "dog": 1}},
		{input: "I ate a donut. Then I ate another donut.", expected: map[string]int{"i": 2, "ate": 2, "a": 1, "donut.": 2, "then": 1, "another": 1}},
		{input: "A man a plan a canal panama.", expected: map[string]int{"a": 3, "man": 1, "plan": 1, "canal": 1, "panama.": 1}},
	}

	for _, tc := range tests {
		m := WordCount(tc.input, wcase.INSENSITIVE)
		assert.NotNil(t, m)
		for i, v := range tc.expected {
			assert.Equal(t, v, m[i], fmt.Sprintf("input: %s", i))
		}
	}
}
