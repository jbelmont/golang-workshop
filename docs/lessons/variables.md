# Variables

Variables in Go can be declared and initialized in several different ways.

Declare variable with `var` keyword
```go
var pizza = "pizza"
```

Notice that we didn't specify a type of `string` it was inferred

Alternatively you can write the statement like this
```go
var pizza string = "pizza"
```

Inside functions, loops, and control flow statements you can use a shorthand operator
```go
func printName(name string) string {
    modName := name + "is a person"
    return modName
}
```

## Naming variables

* Names must start with a letter and may contain letters, numbers or the _ (underscore) symbol