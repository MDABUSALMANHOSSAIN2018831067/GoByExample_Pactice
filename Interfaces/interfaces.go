package main

import "fmt"

type Animal interface {
	walk() string
	eat() string
}

type Dog struct {
	w string
	e string
}

func (d Dog) walk() string {
	return d.w
}

func (d Dog) eat() string {
	return d.e
}

type Cow struct {
	w string
	e string
}

func (c Cow) walk() string {
	return c.w
}

func (c Cow) eat() string {
	return c.e
}

func main() {

	var a Animal = Dog{"Dog is walking.....!!!", "Dog is eating meat....!!!"}
	fmt.Println("Dog......!")
	fmt.Println(a.walk())
	fmt.Println(a.eat())

	fmt.Println()
	fmt.Println("Cow............!!!!!")
	a = Cow{"Cow is walking....!", "Cow is eating grass...!"}

	fmt.Println(a.walk())
	fmt.Println(a.eat())
}
