package utils

import (
	"github.com/hattan/demo/pkg/animals"
)

func GetDogs() (animals.Dogs, error) {
	dogs := make([]*animals.Dog, 0, 5)
	fido, err := animals.NewDog("fido", 1, animals.Corgie)
	if err != nil {
		return nil, err
	}
	dogs = append(dogs, fido)

	bob, err := animals.NewDog("bob", 40, animals.Frenchie)
	if err != nil {
		return dogs, err
	}
	dogs = append(dogs, bob)

	mt, err := animals.NewDog("mt", 3, animals.Mutt)
	if err != nil {
		return dogs, err
	}
	dogs = append(dogs, mt)
	return dogs, nil
}
