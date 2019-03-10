// "multiple-return-values" is a demonstrator for how multiple values can be returned from a function in Go. See https://gobyexample.com/multiple-return-values for more information.

package main

import "fmt"

func CreateGreeting(name string) (primary string, alternate string) {
	primary = "Hello, " + name
	alternate = "Greetings, " + name
	return
}

func main() {
	greeting, alternate := CreateGreeting("Andy")
	fmt.Println(greeting)
	fmt.Println(alternate)
}
