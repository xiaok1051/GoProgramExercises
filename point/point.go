package main

import (
	"fmt"
)

func main() {
	i, j := 45, 1973
	p := &i
	fmt.Println(*p)
	*p = 27
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(*p)
}
