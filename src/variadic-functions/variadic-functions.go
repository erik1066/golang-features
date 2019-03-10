// "variadic-functions" is a demonstrator for how variadic functions work in Go. See https://gobyexample.com/variadic-functions for more information.

package main

import "fmt"

func Greeting(name string, messages ...string) {
	for _, message := range messages {
		fmt.Println(message, name)
	}
}

func main() {
	Greeting("Andy", "Hello", "Greetings", "Salutations")
}
