package main

import (
	"fmt"
)

func main() {
	var a [10]int
	a[0] = 100
	a[1] = 98
	a[3] = 99
	a[4] = 95

	fmt.Println(a[0], a[1], a[2], a[3], a[4])
	fmt.Println(a)

	var b [2]string
	b[0] = "Hello"
	b[1] = "World"
	fmt.Println(b[0], b[1])
	fmt.Println(b)
}
