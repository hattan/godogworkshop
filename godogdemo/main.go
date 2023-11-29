package main

import (
	"fmt"
	"log/slog"

	"github.com/hattan/demo/internal/utils"
)

func main() {
	file := utils.ConfigureLogger("baddirectory/test.log")
	defer func() {
		file.Close()
	}()

	dogs, err := utils.GetDogs()
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
