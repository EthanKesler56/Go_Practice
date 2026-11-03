package main

type Person interface {
	getName() string
	getAge() int
	getStats() int
}

type Hero struct {
	name     string
	age      int
	Strength int
}

type Mage struct {
	name         string
	age          int
	intelligence int
}

func (m Mage) getName() string {
	return m.name
}
func (m Mage) getAge() int {
	return m.age
}
func (m Mage) getStats() int {
	return m.intelligence
}

func (h Hero) getName() string {
	return h.name
}
func (h Hero) getAge() int {
	return h.age
}
func (h Hero) getStats() int {
	return h.Strength
}

func GetInfo(p Person) (string, int) {
	m, ok := p.(Mage)
	if ok {
		intelligence := m.intelligence

		return "This person is a mage!", intelligence
	}
	if !ok {
		intelligence := 9
		return "This person is not a mage!", intelligence
	}
	return "Error!", 0
}

func main() {
	nerissa := Mage{name: "Nerrisa Ravencroft", age: 26, intelligence: 19}
	noel := Hero{name: "Noel Shirogane", age: 26, Strength: 19}
	print(GetInfo(nerissa))
	print(GetInfo(noel))

}
