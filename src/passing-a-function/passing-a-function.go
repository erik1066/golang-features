// "passing-a-function" demonstrates how to pass a function to another function

package main

import "fmt"

func CreateGreeting(name string, generator func(string) string) string {
	greeting := generator(name)
	return greeting
}

func CreateGreeting1(name string) string {
	return "Greetings, " + name + "!"
}

func CreateGreeting2(name string) string {
	return "Hello, " + name
}

func main() {
	greeting1 := CreateGreeting("Andy", CreateGreeting1)
	fmt.Println(greeting1)

	greeting2 := CreateGreeting("Andy", CreateGreeting2)
	fmt.Println(greeting2)
}
