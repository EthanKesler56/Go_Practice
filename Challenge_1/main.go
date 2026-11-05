package main

import "fmt"

func main() {
	nerissa := Mage{name: "Nerrisa Ravencroft", age: 26, intelligence: 19, mana_points: 50}
	GetInfo(nerissa)
	spells := [2]string{"Doom", "Fireball"}
	fmt.Println(getSpellCost(spells))

}
