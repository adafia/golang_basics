package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (Pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (Pointer reciever)
func (p *Person) getMarried(spouceLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouceLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Edem", lastName: "Adafia", city: "Kigali", gender: "m", age: 31}

	// alternative way to init a person struct
	person2 := Person{"Ariel", "Jagusztyn", "Kigali", "f", 32}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.getMarried("Jagusztyn")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
	person2.getMarried("Adafia")
	fmt.Println(person2.greet())
}
