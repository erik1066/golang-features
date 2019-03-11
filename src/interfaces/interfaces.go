// "interfaces" is a demonstrator for how interfaces work in Go. See https://gobyexample.com/interfaces for more information.

package main

import (
	"fmt"
	"strings"
)

type uppercaser struct {
	data string
}

type reverser struct {
	data string
}

type transformer interface {
	transform() string
}

func (t reverser) transform() string {
	runes := []rune(t.data)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (t uppercaser) transform() string {
	var upper = strings.ToUpper(t.data)
	return upper
}

func transformData(t transformer) {
	fmt.Println(t.transform())
}

func main() {
	var reverser = reverser{data: "John"}
	var uppercaser = uppercaser{data: "Dinah"}

	transformData(reverser)
	transformData(uppercaser)
}
