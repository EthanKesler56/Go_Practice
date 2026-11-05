package main

func getSpellCost(spells [2]string) []int {
	spell_cost := make([]int, len(spells))
	for i := 0; i < len(spell_cost); i++ {
		spell_cost[i] = len(spells[i])
	}

	return spell_cost
}
