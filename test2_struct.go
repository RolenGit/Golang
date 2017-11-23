package main

import (
	"fmt"
)

type Myint int

type Book struct {
	teacher string
	age     int
	tall    int
}

func printbook(book Book) {
	fmt.Println("Book's teacher =", book.teacher)
	fmt.Println("Book's agr =", book.age)
	fmt.Println("Book's tall =", book.tall)
}

func printbook1(book1 *Book) {

	fmt.Println("Book's teacher =", book1.teacher)
	fmt.Println("Book's agr =", book1.age)
	fmt.Println("Book's tall =", book1.tall)

}

type Teacher struct {
	name  string
	age   int
	calss string
	zhlei int
}

func (this *Teacher) setTitle(newName string) {
	this.name = newName
}

func (this *Teacher) PrintSelf() {
	fmt.Println("teachr 's name ", this.name)
	fmt.Println("teachr 's class ", this.calss)
	fmt.Println("teachr 's age ", this.age)
	fmt.Println("teachr 's zhlei ", this.zhlei)
}

func main() {
	var a Myint = 10
	fmt.Println(a)

	fmt.Println("-----------------")

	var book1 Book

	book1.age = 37
	book1.tall = 168
	book1.teacher = "榴弹兵"

	printbook(book1)
	fmt.Println("=====================")

	printbook1(&book1)

	var teachr Teacher

	teachr.age = 24
	teachr.calss = "传智播客"
	teachr.name = "安楠楠"
	teachr.zhlei = 407

	fmt.Println("=====================")

	fmt.Println("teachr 's name ", teachr.name)
	fmt.Println("teachr 's class ", teachr.calss)
	fmt.Println("teachr 's age ", teachr.age)
	fmt.Println("teachr 's zhlei ", teachr.zhlei)
	fmt.Println("=====================")

	teachr.name = "王飞"
	//	fmt.Println("name 's ", teachr.setTitle("话下"))
	teachr.setTitle("话下")
	teachr.PrintSelf()

}
