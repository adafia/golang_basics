package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Declare and assign
	vegieArr := [2]string{"Lettuce", "Kale"}

	fmt.Println(vegieArr)


	//Slices
	fruitSlice := []string{"Mango", "Pawpaw", "Grape", "Cherry"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}


