// "methods" is a demonstrator for how methods work in Go. See https://gobyexample.com/methods for more information.

package main

import "fmt"

// Person struct represents a person
type Person struct {
	name    string
	address string
	age     int32
}

// Greet prints a greeting for a person
func (person Person) Greet() {
	fmt.Println("Hello, " + person.name)
	person.name = "Jane" // doesn't change the person's name, as person has been passed by value (copied)
}

// SetName changes the person's name
func (person *Person) SetName(name string) {
	person.name = name
}

func main() {
	person := Person{"John", "123 Birch Street", 45}
	person.Greet()
	fmt.Println(person.name) // still John even though 'Greet' sets name to 'Jane'

	person.SetName("Dinah")
	fmt.Println(person.name) // prints "Dinah"
}
