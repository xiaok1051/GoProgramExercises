package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go!!!")	
		c <- true
		/*使用for range对channel变量进行迭代时，必须进行关闭，否则for range会一直读取chafnel的值，形成死锁*/
		close(c)
	}()
	for v := range c {
		fmt.Println(v)	
	}
}
