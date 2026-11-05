package main

import "fmt"

func main() {
	nerissa := Mage{name: "Nerrisa Ravencroft", age: 26, intelligence: 19, mana_points: 50}
	GetInfo(nerissa)
	num := []int{1, 2, 3, 4}
	fmt.Println(slice_test(num))

}
