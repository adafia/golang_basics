package main

import "fmt"

func main() {
	x := 15
	y := 10

	// if else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}


	// else if
	colour := "green"

	if colour == "red" {
		fmt.Println("colour is red")
	} else if colour == "blue" {
		fmt.Println("colour is blue")
	} else {
		fmt.Println("colour is neither blue nor red")
	}


	// switch
	switch colour {
	case "red":
		fmt.Println("colour is red")
	case "blue":
		fmt.Println("colour is blue")
	default:
		fmt.Println("colour is neither blue nor red")
	}

}


