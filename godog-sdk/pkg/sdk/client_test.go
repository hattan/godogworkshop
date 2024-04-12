package sdk

import (
	"testing"

	animals "github.com/hattan/godog"
	"github.com/stretchr/testify/assert"
)

type Dog = animals.Dog

func TestClientExists(t *testing.T) {
	// arrange
	options := GoDogOptions{
		Host: "localhost",
		Port: "8080",
	}

	// act
	client := NewGoDogClient(options)

	// assert
	assert.NotNil(t, client)
}

func TestCanFetchDogs(t *testing.T) {
	// arrange
	options := GoDogOptions{
		Host: "http://localhost",
		Port: "8080",
	}
	client := NewGoDogClient(options)
	// act
	dogs := client.GetDogs()
	// assert
	assert.IsType(t, []Dog{}, dogs)
	assert.NotNil(t, dogs)
	assert.Equal(t, 2, len(dogs))
}

func TestGetUri(t *testing.T) {
	// arrange
	options := GoDogOptions{
		Host: "http://localhost",
		Port: "8080",
	}
	client := NewGoDogClient(options)
	//act
	uri := client.getUri()
	//assert
	assert.Equal(t, "http://localhost:8080", uri)
}

func TestGetUriWorksWithNoPort(t *testing.T) {
	// arrange
	options := GoDogOptions{
		Host: "https://example.com",
	}
	client := NewGoDogClient(options)
	//act
	uri := client.getUri()
	//assert
	assert.Equal(t, "https://example.com", uri)
}
