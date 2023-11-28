package person

type Person struct {
	FirstName string
	LastName  string
	age       int
}

func NewPerson(firstName, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		age:       0,
	}
}

func (p *Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func (p *Person) GetAge() int {
	return p.age
}

func (p Person) SetFirstNameValue(firstName string) {
	p.FirstName = firstName
}

func (p *Person) SetFirstNamePtr(firstName string) {
	p.FirstName = firstName
}
