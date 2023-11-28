package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/hattan/learngo/pkg/context"
	"github.com/stretchr/testify/assert"
)

func Test_CallApi(t *testing.T) {
	// Arrange
	start := time.Now()
	context := context.NewContext()
	ch, _, _ := context.GetItems()
	expectedTime := 2

	// Act
	CallApi(context, expectedTime)
	result := <-ch
	elapsed := int(time.Since(start).Seconds())

	// Assert
	assert.NotNil(t, result, "Result should not be nil")
	assert.True(t, elapsed >= expectedTime && elapsed <= (expectedTime+1), fmt.Sprintf("Elapsed time should be at least %d seconds", expectedTime))
}
