package main

import "fmt"

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

func GetInfo(p Person) int {
	m, ok := p.(Mage)
	if ok {
		intelligence := m.intelligence
		fmt.Printf("This person has %v int!", intelligence)
		return 1
	}
	if !ok {
		return 2
	}
	return 0
}

func main() {
	nerissa := Mage{name: "Nerrisa Ravencroft", age: 26, intelligence: 9}

	print(GetInfo(nerissa))

}
