package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Func_ReturnsOne(t *testing.T) {
	// Act
	result := add(1, 2)

	// Assert
	assert.True(t, result == 3)
}

func Test_Func_ReturnsMultiple(t *testing.T) {
	// Act
	// AddE only allows positive numbers
	result, err := addE(1, 2)
	if err != nil {
		assert.Fail(t, "Expected Error when passing negative numbers")
		return
	}

	// Assert
	assert.Equal(t, 3, result)
}

func Test_Func_ReturnsMultipleError(t *testing.T) {
	// Act
	// AddE only allows positive numbers
	_, err := addE(-11, 2)
	if err != nil {
		assert.NotNil(t, err)
		return
	}

	// Assert
	assert.Fail(t, "Expected Error when passing negative numbers")
}

func Test_Func_CanBeVariables(t *testing.T) {
	// Arrange
	var f func(int, int) int
	f = add

	// Act
	result := f(1, 2)

	// Assert
	assert.Equal(t, 3, result)
}

func Test_Func_CanBeVariablesShortHand(t *testing.T) {
	// Arrange
	f := add

	// Act
	result := f(10, 2)

	// Assert
	assert.Equal(t, 12, result)
}

func Test_Func_InlineFunctions(t *testing.T) {
	// Arrange
	add2 := func(x, y int) int {
		return x + y
	}

	// Act
	result := add2(10, 2)

	// Assert
	assert.Equal(t, 12, result)
}

func Test_Func_InlineFunctionsImmediatelyInvoked(t *testing.T) {
	// Act
	result := func(x, y int) int {
		return x + y
	}(10, 2)

	// Assert
	assert.Equal(t, 12, result)
}

func Test_Func_ResultsCanBeNamed(t *testing.T) {
	// Arrange
	add := func(x, y int) (result int) {
		result = x + y // notice result is only assigned and not declared (ie = , not :=)
		return result
	}
	// Act

	result := add(10, 2)

	// Assert
	assert.Equal(t, 12, result)
}

func Test_Func_ResultsNakedReturnCanbeUsed(t *testing.T) {
	// Arrange
	add := func(x, y int) (result int) {
		result = x + y
		return
	}
	// Act

	result := add(10, 12)

	// Assert
	assert.Equal(t, 22, result)
}

func Test_Func_VariadacFunctions(t *testing.T) {
	// Arrange
	sum := func(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}

	// Assert
	assert.Equal(t, 3, sum(3))
	assert.Equal(t, 6, sum(1, 2, 3))
	assert.Equal(t, 10, sum(1, 2, 3, 4))
	assert.Equal(t, 20, sum(2, 3, 4, 5, 6))
}

func Test_Func_Closures(t *testing.T) {
	// Arrange
	createFunc := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}

	nextInt := createFunc()

	// Assert
	assert.Equal(t, 1, nextInt())
	assert.Equal(t, 2, nextInt())
	assert.Equal(t, 3, nextInt())
}

func BenchmarkFunc_Closures(b *testing.B) {
	// Arrange
	b.ReportAllocs()
	var local int

	createFunc := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}
	nextInt := createFunc()

	for i := 0; i < b.N; i++ {
		local = nextInt()
	}
	globalValue = local
}

// go test -bench=. -count=3
func BenchmarkSumValue(b *testing.B) {
	b.ReportAllocs()
	var local int
	for i := 0; i < b.N; i++ {
		local = sumValue(i, i)
	}
	globalValue = local
}

func BenchmarkSumPtr(b *testing.B) {
	b.ReportAllocs()
	var local *int
	for i := 0; i < b.N; i++ {
		local = sumPtr(i, i)
	}
	globalValue = *local
}
