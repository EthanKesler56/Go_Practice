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
	mana_points  int
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
