package testruct

import "fmt"

// Member struct
type Member struct {
	id   int
	name string
}

// sayMember
func sayMember() {
	member1 := &Member{1, "pero"}
	member2 := new(Member)
	member2.id = 2
	member2.name = "pero2"
	fmt.Println(member1, member2)
}
