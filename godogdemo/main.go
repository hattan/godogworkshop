package main

import (
	"fmt"
	"log/slog"

	"github.com/hattan/demo/pkg/animals"
)

func getDogs() ([]*animals.Dog, error) {
	dogs := make([]*animals.Dog, 0, 5)
	fido, err := animals.NewDog("fido", 1, animals.Corgie)
	if err != nil {
		return nil, err
	}
	dogs = append(dogs, fido)

	bob, err := animals.NewDog("bob", 2, animals.Frenchie)
	if err != nil {
		slog.Error(err.Error())
		return dogs, err
	}
	dogs = append(dogs, bob)

	mt, err := animals.NewDog("mt", 3, animals.Mutt)
	if err != nil {
		slog.Error(err.Error())
		return dogs, err
	}
	dogs = append(dogs, mt)
	return dogs, nil
}

func main() {
	file := configureLogger("test.log")
	defer func() {
		file.Close()
	}()

	dogs, err := getDogs()
	if err != nil {
		slog.Error(err.Error())
	}

	for _, dog := range dogs {
		if dog.IsMutt() {
			fmt.Println("Mutt", dog.Name)
		}
		dog.Display()
	}

}
