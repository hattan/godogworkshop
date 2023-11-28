package animals

type Animal interface {
	Display()
}

func ShowAnimal(animal Animal) {
	animal.Display()
}
