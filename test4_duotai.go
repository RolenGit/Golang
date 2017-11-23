package main

import (
	"fmt"
)

type Animal interface {
	sleep()
	color() string
	Type() string
}

//定义一个类 只需要重写接口方法就行了
type cat struct {
	color string
}

func (this *cat) sleep() {
	fmt.Println("cat is sleep")
}

func (this *cat) color() string {
	return this.color
}

func (this *cat) Type() string {
	return "cat"
}

//定义一个类 只需要重写接口方法就行了
type dog struct {
	color string
}

func (this *dog) sleep() {
	fmt.Println("cat is sleep")
}

func (this *dog) color() string {
	return this.color
}

func (this *dog) Type() string {
	return "dog"
}

func main() {

	var animal Animal
	animal = &cat{"red"}
	animal = &dog{"green"}

}
