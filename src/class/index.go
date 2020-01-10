package class

// Person 类
type Person struct {
	name string
	age  int
}

// SetName 给person对象设置名字
func (p *Person) SetName(name string) {
	p.name = name
}

// GetName 取person对象名字
func (p *Person) GetName() string {
	return p.name
}
