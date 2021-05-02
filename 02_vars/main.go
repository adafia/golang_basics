package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 - float64
	// complex64 complex128

	// using the var keyword
	var name string = "Edem"
	var age int32 = 31
	const isCool = true

	// short hand (can not live outside of a function)
	surname := "Adafia"
	size := 1.6
	email, height := "adafia.samuel@gmail.com", 5.9

	fmt.Println(name, surname, age, size, isCool, email, height)
	fmt.Printf("%T\n", size)
}