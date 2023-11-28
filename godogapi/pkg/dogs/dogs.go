package dogs

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"

	animals "github.com/hattan/godog"
)

type Dogs []*animals.Dog

var dogs Dogs

func GetDogs() Dogs {
	if dogs == nil {
		initDogs()
	}
	return dogs
}

func initDogs() {
	dogs = make(Dogs, 0)

	fido, _ := animals.NewDog("fido", 2, animals.Frenchie)
	dogs = append(dogs, fido)

	bob, _ := animals.NewDog("bob", 2, animals.Corgie)
	dogs = append(dogs, bob)
}

func GetDogByName(name string) *animals.Dog {
	for _, dog := range dogs {
		if dog.Name == name {
			return dog
		}
	}
	return nil
}

func AddDog(dog *animals.Dog) {
	if dogs == nil {
		initDogs()
	}
	dogs = append(dogs, dog)
}

func RemoveDog(name string) error {
	if dogs == nil {
		initDogs()
	}
	newDogs := make(Dogs, 0)
	found := false
	for _, dog := range dogs {
		if dog.Name != name {
			newDogs = append(newDogs, dog)
		} else {
			found = true
		}
	}

	if found {
		dogs = newDogs
		return nil
	} else {
		message := fmt.Sprintf("delete failed a dog named %s was not found", name)
		slog.Error(message)
		return errors.New(message)
	}
}

func BreedFromString(breed string) (animals.DogBreed, error) {
	switch strings.ToLower(breed) {
	case "corgie":
		return animals.Corgie, nil
	case "frenchie":
		return animals.Frenchie, nil
	case "poodle":
		return animals.Poodle, nil
	case "mutt":
		return animals.Mutt, nil
	case "":
		return -1, errors.New("missing Dog Breed")
	default:
		return -1, errors.New("invalid Dog Breed")
	}
}
