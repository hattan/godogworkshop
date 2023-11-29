package animals

import (
	"fmt"
	"log/slog"
)

type Dog struct {
	Name string
	age  int
	DogBreed
}

type Dogs []*Dog

func NewDog(name string, age int, breed DogBreed) (*Dog, error) {
	if age <= 0 {
		return nil, fmt.Errorf("age must be greater than zero, %s", name)
	}
	if age > 35 {
		return nil, fmt.Errorf("age must be 35 or younger, %s", name)
	}

	dog := Dog{
		Name:     name,
		age:      age,
		DogBreed: breed,
	}

	slog.Info(fmt.Sprintf("New Dog Created %s", dog.Name))
	return &dog, nil
}

func (d *Dog) Display() {
	fmt.Printf("Dog\n Name:%s\n Age:%d\n Breed:%s\n", d.Name, d.age, d.DogBreed)
}

func (d *Dog) SetAge(age int) {
	d.age = age
}
