package testruct

import "fmt"

// Animal
type Animal struct {
	Name string
	Age  int
}

// Cat 组合Animal
type Cat struct {
	A Animal
}

// Dog 匿名组合Animal
type Dog struct {
	Animal
}

// Run is
func (a Animal) Run() {
	fmt.Println(a.Name + "is running")
}
