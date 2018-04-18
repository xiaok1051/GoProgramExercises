//map元素的值类型是一个结构体
package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	m["Bei Jing"] = Vertex{39.26, -115.25}

	fmt.Println(m)
	fmt.Println("----------I'm just a line.----------")

	for lat, _ := range m {
		fmt.Println(lat, ": ", m[lat].Lat, m[lat].Long)
	}
	delete(m, "Bei Jing")
	fmt.Println(m)
}
