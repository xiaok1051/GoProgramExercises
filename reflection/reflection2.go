package main

/*反射嵌入结构*/
import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func main() {
	m := Manager{User: User{1, "OK", 12}, title: "123"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
}
