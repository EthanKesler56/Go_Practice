package main

type getCardInfo interface {
	getName() string
	getColor() string
	getManaCost() int
}

type card struct {
	name      string
	color     string
	mana_cost int
}

func (c card) getName() string {
	return c.name
}

func (c card) getManaCost() int {
	return c.mana_cost
}

func (c card) getColor() string {

	return c.color
}

func main() {
	myCard := card{

		name:      "Ghoulcaller Gisa",
		color:     "Black",
		mana_cost: 4,
	}

	print(getCardInfo.getColor(myCard))

}
