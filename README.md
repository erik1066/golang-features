# Go Language Features

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

```go
const (
    Version = 1
    OS      = "Ubuntu 18.04"
)
```