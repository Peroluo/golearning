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
