package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_basicArray(t *testing.T) {
	// Arrange
	var a [3]int

	// Act
	a[0] = 1
	a[1] = 2
	a[2] = 3

	// Assert
	assert.Equal(t, 3, len(a))
	assert.Equal(t, 1, a[0])
	assert.Equal(t, 2, a[1])
	assert.Equal(t, 3, a[2])
}

func Test_ArrayLiteral(t *testing.T) {
	// Arrange
	a := [3]int{1, 2, 3}

	//Assert
	assert.Equal(t, 2, a[1])
	assert.Equal(t, 1, a[0])
	assert.Equal(t, 2, a[1])
	assert.Equal(t, 3, a[2])
}

func Test_Slice(t *testing.T) {
	// Arrange
	a := [3]int{1, 2, 3}
	//          0  1  2

	// Act
	var s []int = a[0:2]

	// Assert
	assert.Equal(t, 2, len(s))
	assert.Equal(t, 1, s[0])
	assert.Equal(t, 2, s[1])
}

func Test_SliceModification(t *testing.T) {
	// Arrange
	a := [3]int{1, 2, 3}

	// Act
	var s []int = a[1:3]
	s[0] = 5

	//Assert
	assert.Equal(t, 5, a[1])
}

func Test_SliceLiteral(t *testing.T) {
	// Arrange
	s := []int{1, 2, 3}

	//Assert
	assert.Equal(t, 3, len(s))
}

func Test_SliceBounds(t *testing.T) {
	// Arrange
	s := []int{1, 2, 3, 4, 5, 6}
	s = s[:2]

	// Assert
	assert.Equal(t, 2, len(s))
}

func Test_SliceCapVsLength(t *testing.T) {
	// Arrange
	a := [6]int{1, 2, 3, 4, 5, 6}
	//          0  1  2  3  4  5

	// Act
	s := a[1:4]

	// Assert

	assert.Equal(t, 3, len(s))
	assert.Equal(t, 5, cap(s))
}

func Test_SliceCreationWithMake(t *testing.T) {
	// Arrange
	a := make([]int, 5)

	// Assert
	assert.Equal(t, 0, a[0])
	assert.Equal(t, 5, len(a))
	assert.Equal(t, 5, cap(a))
}

func Test_SliceCreationWithMakeWithCap(t *testing.T) {
	// Arrange
	a := make([]int, 5, 10)

	// Assert
	assert.Equal(t, 0, a[0])
	assert.Equal(t, 5, len(a), "len")
	assert.Equal(t, 10, cap(a), "cap")
}

func Test_SliceAppend(t *testing.T) {
	// Arrange
	a := make([]int, 5)

	// Act
	for i := 0; i < 7; i++ {
		a = append(a, 100*i)
	}

	// Assert
	assert.Equal(t, 100, a[5])

	assert.Equal(t, 6, len(a), "len")
	assert.Equal(t, 10, cap(a), "cap")
}

func Test_SliceRange(t *testing.T) {
	// Arrange
	a := []int{100, 200, 300, 800, 1600, 3200, 6400, 12800, 25600, 51200}

	// Act
	sum := 0
	for _, value := range a {
		sum = sum + value
	}

	// Assert
	assert.Equal(t, 102200, sum)
}
