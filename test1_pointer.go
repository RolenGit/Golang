package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	var p *int

	if p == nil {
		fmt.Println("p is nil")
	}
	//	fmt.Println(*p)
	//	fmt.Println(p)

	a := []int{19, 28, 87}

	var ptr [3]*int

	fmt.Println("----------------")

	for i := 0; i < 3; i++ {
		ptr[i] = &a[i]

	}

	for _, value := range a {
		fmt.Println(value)
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
