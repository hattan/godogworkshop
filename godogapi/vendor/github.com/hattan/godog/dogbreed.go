package animals

type DogBreed int

const (
	Corgie   DogBreed = 0
	Frenchie DogBreed = 1
	Poodle   DogBreed = 2
	Mutt     DogBreed = 3
)

func (breed DogBreed) String() string {
	switch breed {
	case Corgie:
		return "Corgie"
	case Frenchie:
		return "Frenchie"
	case Poodle:
		return "Poodle"
	case Mutt:
		return "Mutt"
	default:
		return "unknown"
	}
}

func (breed DogBreed) IsMutt() bool {
	return breed == Mutt
}
