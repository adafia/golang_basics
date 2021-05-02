package main

import "fmt"

func main(){
	// Define a map
	emails := make(map[string]string)

	// assing key-value pairs
	emails["Edem"] = "edem@gmail.com"
	emails["Ariel"] = "ariel@gmail.com"
	emails["Sambs"] = "sambleen@gmail.com"


	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Edem"])

	// Delete from map
	delete(emails, "Sambs")
	fmt.Println(emails)


	// declare and assign key-value pairs
	contacts := map[string]string{"Edem": "0783564602", "Ariel": "0783564603"}
	contacts["Sambs"] = "0783564602"

	fmt.Println(contacts)
	fmt.Println(len(contacts))
	fmt.Println(contacts["Edem"])
}