package main

import "fmt"

func main() {
	a := 5

	// b is assigned the memory address of a
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// Use * to read val from memory address
	fmt.Println(*b)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}


