// "ranges" is a demonstrator for how ranges work in Go. See https://gobyexample.com/range for more information.

package main

import "fmt"

func main() {

	// create a slice
	places := make([]string, 5)

	// add values to the slice
	places[0] = "New York City"
	places[1] = "Los Angeles"
	places[2] = "Chicago"
	places[3] = "Toronto"
	places[4] = "Victoria"

	// append a new value, increasing slice size to 5
	places = append(places, "Seattle")

	// output all of the items in the slice
	fmt.Println("## Print all the places:")
	for i := 0; i < len(places); i++ {
		fmt.Println(places[i])
	}

	// get a range from the slice and assign it to a new slice
	canadianPlaces := places[3:5]

	// output the Canadian places
	fmt.Println("## Print Canadian places:")
	for i := 0; i < len(canadianPlaces); i++ {
		fmt.Println(canadianPlaces[i])
	}

	// output big cities only, this time using for loop syntax
	fmt.Println("## Print big U.S. cities:")
	for i, place := range places[:3] {
		fmt.Println(i+1, place)
	}

	// get all elements in the slice after the 4th
	otherPlaces := places[5:]
	fmt.Println("## Print other cities:")
	for _, place := range otherPlaces {
		fmt.Println(place)
	}
}
