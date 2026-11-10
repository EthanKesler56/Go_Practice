package main

import (
	"fmt"
)

func send(ch chan int) {
	ch <- 99
}

func main() {
	ch := make(chan int)
	go send(ch)
	fmt.Println(<-ch)
}
