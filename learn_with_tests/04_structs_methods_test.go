package test

import (
	"fmt"
	"testing"

	"github.com/hattan/learngo/pkg/drawer"
	"github.com/hattan/learngo/pkg/foo"
	"github.com/hattan/learngo/pkg/person"
	"github.com/stretchr/testify/assert"
)

func Test_Method_TypeAlias(t *testing.T) {
	// Arrange
	type KeyName string
	test := func(k KeyName) string {
		return fmt.Sprintf("k=%s", k)
	}

	// Act
	var k KeyName = "test"

	// Assert
	assert.Equal(t, "k=test", test(k))
}

func Test_Method_Struct(t *testing.T) {
	// Act
	bob := person.Person{
		FirstName: "Bob",
		LastName:  "NotSmith",
	}

	// Assert
	assert.NotNil(t, bob)
	assert.Equal(t, "Bob", bob.FirstName)
	assert.Equal(t, "NotSmith", bob.LastName)
}

func Test_Method_StructNewPattern(t *testing.T) {
	// Act
	airplane := person.NewPerson("Airplane", "McAirPlaneFace")

	// Assert
	assert.NotNil(t, airplane)
	assert.Equal(t, "Airplane", airplane.FirstName)
	assert.Equal(t, "McAirPlaneFace", airplane.LastName)
}

func Test_Method_StructWithFullName(t *testing.T) {
	// Arrange
	boyScout := person.NewPerson("Christopher", "Pike")

	// Act
	fullName := boyScout.FullName()

	// Assert
	assert.NotNil(t, fullName)
	assert.Equal(t, "Christopher Pike", fullName)
}

func Test_Method_StructNotExportedField(t *testing.T) {
	// Arrange
	bob := person.Person{
		FirstName: "Bob",
		LastName:  "NotSmith",
		// age:       0, // cannot set age since it is not exported
	}

	// Act
	bob.SetAge(42)

	// Assert
	assert.NotNil(t, bob)
	assert.Equal(t, 42, bob.GetAge())
}

func Test_Method_CanBeDefinedOnNonStructs(t *testing.T) {
	// Arrange
	d := drawer.Drawer("")

	// Act
	line := d.GetLine()

	// Assert
	assert.Equal(t, "---------------------", line)
}

func Test_Method_ValueReceiver(t *testing.T) {
	// Arrange
	vulcan := person.NewPerson("Spock", "")

	// Act
	vulcan.SetFirstNameValue("pointy ears")

	// Assert
	assert.Equal(t, "pointy ears", vulcan.FirstName)
}

func Test_Method_PointerReceiver(t *testing.T) {
	// Arrange
	vulcan := person.NewPerson("Spock", "")

	// Act
	vulcan.SetFirstNamePtr("pointy ears")

	// Assert
	assert.Equal(t, "pointy ears", vulcan.FirstName)
}

func Test_Struct_EmbeddedFields(t *testing.T) {
	// Arrange
	f := foo.Foo{}
	// Act

	f.Count = 42

	// Assert
	assert.Equal(t, 42, f.Bar.Count)
}

func Test_Struct_Promoted_Methods(t *testing.T) {
	// Arrange
	f := foo.Foo{}
	f.Count = 42

	// Assert
	assert.Equal(t, 84, f.Double())
}
