# Go Language Features

[Go by Example](https://gobyexample.com/) is an excellent syntax reference. If you're new to Go, check out the free online book _[An Introduction to Programming in Go](http://www.golang-book.com/books/intro)_. 

See [Pop!_OS setup guide](https://github.com/erik1066/pop-os-setup) for instructions on installing Go on Ubuntu 18.04.

Example "Hello, World!" application in GoLang:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

> "fmt" is a package in Go. It's short for "format".

## Types

There are just a few basic types:

* `bool`
* `string`
* `int` (size depends on 32- or 64-bit arch), `int8`, `int16`, `int32`, `int64`
* `uint` (size depends on 32- or 64-bit arch), `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
* `byte` (alias for `uint8`)
* `rune` (alias for `int32`), basically meant as a `char` like in C++, C#, and Java
* `float32`, `float64`
* `complex64`, `complex128`

> Integer types are incompatible without explicit conversions. For example, an `int32` is incompatible with `int64` and vice-versa.

There are a few advanced types:

* Array
* Slice (like a `List` in C# and Java)
* Struct
* Pointer
* Function
* Interface
* Map (like a `Dictionary` in C# and `Map` in Java)
* Channel

## Variable Assignments

[Example code](src/variables/variables.go)

Go has three methods for declaring and/or assigning variables. The first is to declare a variable using `var`, provide a variable name, define the type, an assignment operator, and then optionally the value you want to assign after the operator. For example:

```go
var message string = "Hello, Go!"
```

The above is equivalent to:

```go
var message string
message = "Hello, Go!"
```

You can optionally omit the `string` from the first example because the right-hand side of the `=` operator implies this is a `string`. Go is intelligent enough to figure out the type based on the right-hand side of the operator, much like C# and Java. Hence the following example is legal:

```go
var message = "Hello, Go!"
```

You can alternatively delcare a variable and assign it without `var` by using the `:=` operator:

```go
message := "Hello, Go!"
```

Declaring and assigning multiple variables at the same time is possible:

```go
var a, b, c = 1, 2, 3
```

## Pointers

[Example code](src/pointers/pointers.go)

Pointers are declared using the `*` operator and dereferenced using the `&` operator.

```go
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
    fmt.Println(name)  // prints "Mary"
}
```

## User-defined types

[Example code](src/user-defined-types/user-defined-types.go)

Go has no concept of classes. Instead, class-like "structures" are created using the `type` keyword with the `struct` type:

```go
type Person struct {
    name string
    age  int
}
```

A structure can be declared and assigned like such:

```go
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
}
```

Note that we can alternatively use the following two syntax options to assign values to a user-defined type:

```go
var sandra = Person{"Sandra", 55}
fmt.Println(sandra) // prints "{Sandra 55}"

var maria = Person{age: 65, name: "Maria"}
fmt.Println(maria) // prints "{Maria 65}"
```

## Constants

[Example code](src/constants/constants.go)

```go
const (
    Version = 1
    OS      = "Ubuntu 18.04"
)
```

## Loops

[Example code](src/loops/loops.go)

There's only one loop syntax in Go: `for`. There is no `do`, `while`, or `foreach`. The `for` syntax can take several forms:

```go
// traditional for loop
for initialization; condition; after {
    // statements
}

// equivalent to a while loop
for condition {
    // statements
}

// equivalent to an infinite loop
for {
    // statements
}
```

Example of the the traditional `for`:

```go
package main

import "fmt"

func main() {
	names := [3]string{"John", "Mary", "Susan"} // creates an array

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
```

Loops in Go can be terminated using `break` or `return`.

## Ranges

[Example code](src/ranges/ranges.go)

A subset of a collection can be obtained using the range sytnax `[start..end]`. For example, the loop below prints just the Canadian city names "Toronto" and "Victoria":

```go
places := make([]string, 6) // create a slice
places[0] = "New York City"
places[1] = "Los Angeles"
places[2] = "Chicago"
places[3] = "Toronto"
places[4] = "Victoria"
places[5] = "Seattle"

canadianPlaces := places[3:5]

for i := 0; i < len(canadianPlaces); i++ {
    fmt.Println(canadianPlaces[i])
}
```

We can also just print the big U.S. cities in the list by using special `for` syntax as shown below. Notice the abscence of the starting value for the range. Leaving the start value unspecified is shorthand for "start at the beginning".

```go
for i, place := range places[:3] {
    fmt.Println(i+1, place)
}
```

In the above example, `i` is the index value and `place` is the value of the slice at that index. Both variables can then be used in the loop body.

A problem arises when we do not need to use `i` in the loop body: Go will refuse to compile our code if there is an unused variable. In those situations we can use the `_` operator in place of `i` to avoid a compile-time error:

```go
otherPlaces := places[5:]
for _, place := range otherPlaces {
    fmt.Println(place)
}
```

Leaving the end value of the rage unspecified, as shown above, is shorthand for "take all the remaining elements".

## Variadic functions

[Example code](src/variadic-functions/variadic-functions.go)

A "variadic function" is one that accepts a variable number of parameters of a specific type. Use the `...` operator to specify that a method parameter is variadic:

```go
func Greeting(name string, messages ...string) {
	for _, message := range messages {
		fmt.Println(message, name)
	}
}

func main() {
	Greeting("Andy", "Hello", "Greetings", "Salutations")
}
```

Output:

```
Hello Andy
Greetings Andy
Salutations Andy
```

## Multiple return values

[Example code](src/multiple-return-values/multiple-return-values.go)

Go can return multiple named values from a function:

```go
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
```

Output:

```
Hello, Andy
Greetings, Andy
```

In the above example, `primary` and `alternate` are the named return values. Observe that they're wrapped in parenthesis after the method signature and before the opening brace.

## Passing functions to functions

[Example code](src/passing-a-function/passing-a-function.go)

Go allows passing a function to another function. Passing functions is a matter of declaring the function-to-be-passed in the receiving function's signature. We can see this in the example below where how the `generator` parameter is declared as a function that accepts a `string` and returns a `string`. We can then supply any function that meets those requirements to `CreateGreeting`.

```go
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
```

Output:

```
Greetings, Andy!
Hello, Andy
```

## Methods

[Example code](src/methods/methods.go)

Methods operate on defined `struct` types. They are implemented by specifying the type in parenthesis prior to the name of the method. For example:

```go
type Person struct {
	name    string
}

func (person Person) Greet() {
	fmt.Println("Hello, " + person.name)
}

func main() {
	person := Person{"John"}
	person.Greet()
}
```

Observe that `Greet()` is called on an instance of `Person` using the `.` operator. `Greet()` also has `(person Person)` prior to the method name, specifying that this method only operates on `Person`.

However, the `Greet()` method will not be able to modify any values on the `person` instance that was passed to it, as `person` was passed by value and not by reference. To modify the instance being operated on we can use a pointer:

```go
type Person struct {
	name    string
}

func (person *Person) SetName(name string) {
	person.name = name
}

func main() {
	person := Person{"John"}
	person.SetName("Dinah")
	fmt.Println(person.name) // prints "Dinah"
}
```

## Interfaces

[Example code](src/interfaces/interfaces.go)

Any type that has the same methods as an interface implements that interface. We can then use that type anywhere that interface is required.

In the example below, the interface `transformer` is defined as having a method called `transform` that outputs a string. Both `reverser` and `uppercaser` implement the `transformer` interface by conversion, since they each have a method that meets the interface definition. We can then use an instance of either `reverser` or `uppercaser` as an argument to the `transformData` function, which expects an interface.

```go
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

	transformData(reverser) // outputs "nhoJ"
	transformData(uppercaser) // outputs "DINAH"
}
```

