package function

// Add 两数相加
func Add(x int, y int) int {
	return x + y
}

// Add2 同参简写
func Add2(x, y int) int {
	return x + y
}

// Swap 多参返回简写
func Swap(a, b string) (x, y string) {
	x = b
	y = a
	return
}
