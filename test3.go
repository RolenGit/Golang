package main

import (
	"fmt"
)

type humen struct {
	name string
	age  int
}

func (this *humen) eat() {
	fmt.Println("humen.eat :", this)
}
func (this *humen) walk() {
	fmt.Println("humen.walk :", this)
}

type superman struct {
	humen
	name  string
	level int
}

func (this *superman) eat() {
	fmt.Println("superman . eat", this)
}

func (this *superman) walk() {
	fmt.Println("superman . walk", this)
}

func test(h *humen) {
	h.eat()
	fmt.Println("test ...", h)
}

func main() {

	h := humen{"wawa", 34}
	h.eat()
	h.walk()

	fmt.Printf("%+v\n", h)

	fmt.Printf("+++++++++++++++++++++++\n")

	super := superman{humen{"傻逼", 34}, "sdasd", 78}

	super.eat()
	super.walk()
	super.humen.walk()

}
