package main

import (
	"fmt"
)

func myFunction(fName string, lName string) (string, error) {
	fullName := fName + " " + lName
	return fullName, nil
}

func addOne() func() int {
	var x int

	return func() int {
		x++
		return x + 1
	}
}

func twoNumberSum(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func main() {
	receiveFullName, s := myFunction("MD. ABU SALMAN", "HOSSAIN")

	if s != nil {
		fmt.Println("Ftal Error: ", s)
	}

	fmt.Println(receiveFullName)

	f := addOne()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("Sumation of two numbers is ", twoNumberSum(2, 3))
	fmt.Println("Sumation of two numbers is ", twoNumberSum(2, 99))
	fmt.Println("Sumation of two numbers is ", twoNumberSum(2, 77))
	fmt.Println("Sumation of two numbers is ", twoNumberSum(2, 65434))
}
