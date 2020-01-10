package point

import "fmt"

// Point 指针
func Point() {
	i := 2
	p := &i
	fmt.Println(*p)
}
