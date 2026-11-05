package main

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
