package main

import (
	"fmt"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vertex) Scale() float64 {
	return v.X + v.Y
}

func main() {
	var a Abser
	f := Vertex{3, 4}
	a = &f

	fmt.Println(f.Abs())
	fmt.Println(f.Scale())
	fmt.Println(a.Abs())
	//fmt.Println(f.Scale())
}
