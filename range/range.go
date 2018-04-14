package main

import (
	"fmt"
)

func main() {
	//range遍历切片的所有元素，并求和
	fibo := []int{2, 3, 5, 8, 13, 34, 55}
	sum := 0
	for _, num := range fibo {
		sum += num
		fmt.Println(num, sum)
	}
	fmt.Println("sum:", sum)

	//range遍历字符串
	for i, c := range "go" {
		fmt.Printf("%d, %c\n", i, c)
	}

	//range遍历字典，返回键值对
	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cofox"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
