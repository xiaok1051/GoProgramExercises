package main

import (
	"fmt"
)

func main() {
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true
}
