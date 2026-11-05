package main

func loops(m Mage) int {
	var x int

	for x := 0; x <= m.mana_points; x++ {
		print("fireball!")
	}
	return x
}
