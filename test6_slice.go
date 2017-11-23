package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3}
	for index, value := range a {
		fmt.Printf("index is :%d, value is %d\n", index, value)
	}
	fmt.Println("===================")

	var b = make([]int, 5)
	b[0] = 10
	b[1] = 13
	b[2] = 14
	b[3] = 15
	b[4] = 16

	for i := 0; i < len(b); i++ {
		fmt.Printf("%d\n", b[i])
	}

	fmt.Println("===================")
	s1 := a[0:2]
	for index, value := range s1 {
		fmt.Printf("index is :%d, value is %d\n", index, value)
	}

	fmt.Println("===================")

	new := make([]int, 6)
	copy(new, b)
	fmt.Printf("new  = %+v\n", new)

}
