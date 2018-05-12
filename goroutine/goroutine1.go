/*chanal最基本的操作*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true
	}()
	<-c
}
