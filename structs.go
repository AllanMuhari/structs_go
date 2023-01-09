package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var per1 Person
	var per2 Person

	// Per1 specification
	per1.name = "Allan"
	per1.age = 45
	per1.job = "Developer"
	per1.salary = 6000

	// Per2 specification
	per2.name = "Mercy"
	per2.age = 24
	per2.job = "Marketing"
	per2.salary = 4500

	// Print Per1 info by calling a function
	printPerson(per1)

	// Print Per2 info by calling a function
	printPerson(per2)
}

func printPerson(per Person) {
	fmt.Println("Name: ", per.name)
	fmt.Println("Age: ", per.age)
	fmt.Println("Job: ", per.job)
	fmt.Println("Salary: ", per.salary)
}
