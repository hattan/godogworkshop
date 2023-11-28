package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewContext_HasItems(t *testing.T) {
	// Arrange
	ctx := NewContext()

	// Act
	ch, done, wg := ctx.GetItems()

	//Assert
	assert.NotNil(t, ch, "Channel should not be nil")
	assert.NotNil(t, done, "Done Channel should not be nil")
	assert.NotNil(t, wg, "Wait Group should not be nil")
}

func Test_NewContext_InnerContext_NotNil(t *testing.T) {
	// Arrange
	ctx := NewContext()

	//Assert
	assert.NotNil(t, ctx.ctx, "Inner context should not be nil")
}

func Test_Channel(t *testing.T) {
	// Arrange
	message := "Hello"
	ctx := NewContext()

	// Act
	ch, _, _ := ctx.GetItems()
	go func() {
		ch <- message
	}()

	//Assert
	result := <-ch
	assert.Equal(t, message, result, "Channel is not sending messages")
}

func Test_DoneChannel(t *testing.T) {
	// Arrange
	ctx := NewContext()

	// Act
	_, done, _ := ctx.GetItems()
	go func() {
		done <- true
	}()

	//Assert
	result := <-done
	assert.Equal(t, true, result, "Done Channel is not sending messages")
}

func Test_WaitGroup(t *testing.T) {
	// Arrange
	ctx := NewContext()
	complete := false

	// Act
	_, _, wg := ctx.GetItems()
	wg.Add(1)

	go func() {
		wg.Done()
		complete = true
	}()

	//Assert
	wg.Wait()
	assert.True(t, complete, "Wait Group is not working")
}
