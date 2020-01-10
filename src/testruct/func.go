package testruct

// Person struct
type Person struct {
	name string
	age  int
}

// NewPerson 构造函数
func NewPerson(name1 string, age1 int) *Person {
	return &Person{name: name1, age: age1}
}

// SetName ..
func (person *Person) SetName(name string) {
	person.name = name
}

// GetName ..
func (person *Person) GetName() string {
	return person.name
}
