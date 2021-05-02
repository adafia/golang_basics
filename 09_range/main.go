package main

import "fmt"

func main(){
	ids := []int{33, 76, 54, 23, 11, 2}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// loop through ids (without using the index)
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)


	// Range with map
	contacts := map[string]string{"Edem": "0783564602", "Ariel": "0783564603"}

	for k, v := range contacts {
		fmt.Printf("%s: %s\n", k, v)
	}

	// loop through keys
	for k := range contacts {
		fmt.Println("Key: "+ k)
	}

}