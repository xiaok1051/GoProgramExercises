package main

import (
	"fmt"
)

type Coodinate struct {
	X int
	Y int
}

var (
	Co1 = Coodinate{73, 45}
	Co2 = Coodinate{X: 75}
	Co3 = Coodinate{}
	p   = &Coodinate{9, 17}
)

func main() {
	fmt.Println(Co1, Co2, Co3, p, *p)

}
