package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["Anna"] = 16
	m["Tom"] = 15
	m["Smith"] = 17

	fmt.Println(m)

	//遍历map
	for k, v := range m {
		if v == 17 {
			fmt.Print("* -->")
		}
		fmt.Println(k, v)
	}

	m["Anna"] = 17
	fmt.Println(m)

	kv, ok := m["anna"]
	fmt.Println(kv, ok)
	kv1, ok1 := m["Anna"]
	fmt.Println(kv1, ok1)

	delete(m, "Anna")
	fmt.Println(m)
}
