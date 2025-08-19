package basics

import "fmt"

var middleName string = "John" // We can declare a variable outside of a function with var keyword

func dataVariables() {
	var age int = 20
	var name string = "John"
	var isStudent bool = true

	fmt.Println(age, name, isStudent)

	age = 30
	name = "Jane"
	isStudent = false

	fmt.Println(age, name, isStudent)

	count := 10
	stringNames := "John"

	fmt.Println(middleName)

	//middleName := "John"
	fmt.Println(count, stringNames)

	/*
		Default values
		Numberic Types: 0
		Boolean Types: false
		String Types: ""
		Pointers , Maps , Slices , Channels s, Functions : nil
	*/

	printName()

}

func printName() {
	firstName := "Michael"
	lastName := "Jackson"

	fmt.Println(firstName, lastName)
}
