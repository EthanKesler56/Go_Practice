package main

import "fmt"

func mapPractice() {
	ages := make(map[string]int)
	ages["Noel"] = 26
	ages["Nerissa"] = 27
	ages["Mori"] = 28
	fmt.Println(ages)

}
func mapStruct() map[string]Mage {
	new_mem := Mage{name: "Kiara", age: 5000, intelligence: 16, mana_points: 90}
	new_mem_two := Mage{name: "Lamy", age: 900, intelligence: 90, mana_points: 100}
	new_mem_three := Mage{name: "Flare", age: 25, intelligence: 91, mana_points: 110}

	members := map[string]Mage{"Test1": new_mem, "Test2": new_mem_two}
	members["Test3"] = new_mem_three
	elem, ok := members["Test2"]
	if ok {
		fmt.Println("Works!", elem)
	}

	return members

}
