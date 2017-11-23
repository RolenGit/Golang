package main

import (
	"fmt"
	"reflect"
)

func Myfunc(a interface{}) string {

	fmt.Println("Myfunc ", a)
	fmt.Println("type is ", reflect.TypeOf(a))

	value, ok := a.(string)

	if !ok {
		fmt.Println("a is not Ok, is not string type")
	}
	fmt.Println("a is ok,is string type")
	return value
}

func main() {
	var value1 int = 100

	Myfunc(value1)

	var value2 string = "abdc"

	Myfunc(value2)
}
