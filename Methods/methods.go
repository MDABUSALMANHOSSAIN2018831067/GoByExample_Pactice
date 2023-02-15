package main

import "fmt"

type Employee struct {
	name, designation string
}

func (e *Employee) updateProfile(newName, newDesignation string) {
	e.name = newName
	e.designation = newDesignation
}

func (e *Employee) printUpdateProfile() {
	fmt.Println("Name : ",e.name)
	fmt.Println("Designation : ",e.designation)
}
// func s(){
// 	fmt.Println("s")
// }

func main(){
	employee := Employee{"MD. ABU SALMAN HOSSAIN","ASSOCIATE SOFTWARE ENGINEER"}
	employee.printUpdateProfile()

	employee.updateProfile("ASHIQUL HASAN SHAKIL","INTERN SOFTWARE ENGINEER")
	employee.printUpdateProfile()

	
}