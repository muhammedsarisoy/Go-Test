package basics

import "fmt"

type EmployeeGoogle struct {
	firstName string
	lastName  string
	age       int
}

type EmployeeMicrosoft struct {
	firstName string
	lastName  string
	age       int
}

func namingConventions() {
	// PascalCase for package names
	// Structs , Interfaces , Enums , Constants , Types , Methods , Functions : PascalCase

	// Snake Case
	// Variables , Functions , Methods , Packages : snake_case

	// Camel Case
	// Variables , Functions , Methods , Packages : camelCase

	// Kebab Case
	// Variables , Functions , Methods , Packages : kebab-case

	const MAXMETRIES = 5

	var employeeId = 1001

	fmt.Println("Employee ID: ", employeeId)

}
