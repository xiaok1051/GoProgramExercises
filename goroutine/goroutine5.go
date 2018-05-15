/*多于一个channe时的处理demo,随机打印0/1*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func(c) {
		for v := range {
			fmt.Pritln(c)
		}	
	}()

	for {
		select {
			case c <- 0
			case c <- 1 
		}
	}
}
