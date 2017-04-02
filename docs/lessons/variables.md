# Variables

[Variables Go Specification](https://golang.org/ref/spec#Variables)

Variables in Go can be declared and initialized in several different ways.

Declare variable with `var` keyword
```go
var pizza string
pizza = "pizza"
```

Alternatively you can write the statement like this
```go
var pizza string = "pizza"
```

Notice that we didn't specify a type of `string` it was inferred

Inside functions, loops, and control flow statements you can use a shorthand operator
```go
func printName(name string) string {
    modName := name + "is a person"
    return modName
}
```

**You cannot use the shorthand operator `:=` outside of this context**
```go
package main

import (
	"fmt"
)

str := "Hello, playground"

func main() {
	fmt.Println(str)
}
```

If you were to try to run this hello world program you would see the following error message
`syntax error: non-declaration statement outside function body`

## Naming variables

* Names must start with a letter and may contain letters, numbers or the _ (underscore) symbol

## Default values

Variables are given default values when they are declared

```go
var num int // 0

var dec float64 // 0

var str string // ""

var isVal bool // false

var num2 *int // nil

type Data struct{} // nil

func someFunc() {} // nil
```

[Variables Playground](https://play.golang.org/p/ZwaSvq5oc2)