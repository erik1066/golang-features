package main

import "fmt"

func main() {
	name := "John"
	var message *string = &name
	fmt.Println(message)  // prints a memory address
	fmt.Println(*message) // prints "John"

	/* Both name and message point to the same memory location. Changing one
	   therefore changes the other. For instance: */

	*message = "Mary"
	fmt.Println(name) // prints "Mary"
}
