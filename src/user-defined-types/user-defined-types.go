package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var bob = Person{}
	bob.name = "Bob"
	bob.age = 45
	fmt.Println(bob) // prints "{Bob 45}"

	var sandra = Person{"Sandra", 55}
	fmt.Println(sandra) // prints "{Sandra 55}"

	var maria = Person{age: 65, name: "Maria"}
	fmt.Println(maria) // prints "{Maria 65}"
}
