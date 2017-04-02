# Conditionals

[If Go Specification](https://golang.org/ref/spec#If_statements)

Golang has control structures as you may have seen them in other programming languages

You just won't see parenthesis `()` around them

```go
value2 := 7
if value2 > 10 {
    value2++
}
```

There exists `if`, `else if`, `else`, `switch` but there is not a ternary operator `?:`

```go
const movie = "Rocky"
var score int
if movie == "Rambo" {
    score = 9
} else if movie == "Rocky" {
    score = 10
} else {
    score = 6
}
```

Traditionally with a ternary operator you could do the following in one line

```go
num := 3
isFlag := num == 3 ? "I am 3" : "Nope"
```

You can do the above statement as follows

```go
var str string
if num == 3 {
    str = "I am 3"
} else {
    str = "Nope"
} 
```

There are some nice features to the `if` statement

```go
if ok := strings.Contains(val, "foo"); ok {
    return true
}
return false
```

Notice here we assign a variable named ok and then if ok is true it will enter the `if` block

#### Switch statement

Golang also has a `switch` statement which can be used like this

[Switch Go Specification](https://golang.org/ref/spec#Switch_statements)

```go
var food string = "lemons"

switch food {
case "apples"
    fmt.Println("yay apples")
case "lemons"
    fmt.Println("Why lemons?")
default:
    fmt.Println("Hmm")
}
```

Notice that the switch statement doesn't have any break statements

* A case body breaks automatically, unless it ends with a fallthrough statement.
* Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
* Switch without a condition is the same as switch true.
* This construct can be a clean way to write long if-then-else chains.

[Conditionals Playground](https://play.golang.org/p/ojJwDs0k99)