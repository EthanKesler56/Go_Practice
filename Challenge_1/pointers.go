package main

import "fmt"

type Mage struct {
	name         string
	intelligence int
}

func pointerPractice() {
	p := 10
	test := &p
	print(test)
	print("HERE", *test)

}

func emptyPointerPractice() {
	var p *int
	fmt.Printf("value of p: %v\n", p)
}

// PassByReference
func increment(x int) {
	x++
	fmt.Println(x)
}

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
}
