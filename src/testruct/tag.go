package testruct

// Teacher struct
type Teacher struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Goot struct
type Goot struct {
	IsGood  bool    `json:"isGood"`
	Teacher Teacher `json:"teacher"`
}

// func main() {
// 	stu := s.Teacher{Name: "wd", Age: 22}
// 	data, err := json.Marshal(stu)
// 	if err != nil {
// 		fmt.Println("json encode failed err:", err)
// 		return
// 	}
// 	fmt.Println(string(data)) //{"name":"wd","age":22}
// }
