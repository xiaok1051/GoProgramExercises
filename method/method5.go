package main

import (
	"fmt"
)

type TZ int

func main() {
	var a TZ
	a.Print()
	(*TZ).Print(&a)
}

func (a *TZ) Print() {
	fmt.Println("TZ")
}
