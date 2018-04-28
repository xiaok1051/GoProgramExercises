package main

import (
	"fmt"
)

//定义一个接口
type Stringer interface {
	String() string
}

//定义一个结构体
type Cofox struct {
	name string
}

//定义一个方法
func (c *Cofox) String() string {
	return "Joel " + c.name
}

func main() {
	var S Stringer
	c := Cofox{"Smith"}
	S = &c

	fmt.Println(S.String())
	fmt.Println(c.String())
	fmt.Println(c.name)
}
