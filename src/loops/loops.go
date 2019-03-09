// "loops" is a demonstrator for how loops work in Go

package main

import "fmt"

func main() {

	// print out a bunch of names from an array (see https://gobyexample.com/arrays for more info on arrays)
	fmt.Println("Names:")

	names := [3]string{"John", "Mary", "Susan"} // creates an array

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// print out a bunch of placenames from a slice (see https://gobyexample.com/slices for more info on slices)
	fmt.Println(" ")
	fmt.Println("Places:")

	places := make([]string, 3) // creates a slice
	places[0] = "New York City"
	places[1] = "Los Angeles"
	places[2] = "Chicago"
	places = append(places, "Seattle")

	for i := 0; i < len(places); i++ {
		fmt.Println(places[i])
	}
}
