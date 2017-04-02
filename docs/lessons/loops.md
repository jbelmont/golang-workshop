# Loops

In Golang there is only one syntax for doing iteration which is done with the `for` loop
At first this may seem like a limitation but the `for` is very powerful in Go

A "for" statement specifies repeated execution of a block. 

There are three forms:

1. The iteration may be controlled by a single condition
2. A "for" clause
3. A "range" clause.

The basic for loop has three components separated by semicolons:

1. An initialization statement: executed before the first iteration
2. A condition expression: evaluated before every iteration
3. Post/Update statement: executed at the end of every iteration

[For Go Specification](https://golang.org/ref/spec#For_statements)

```go
numbers := []int{1,2,3,4,5,6}
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
```

This is very much like a regular `for` loop that you have seen in other languages except it has no parenthesis

*The init and post statement are optional.*

```go
numbers := []int{1, 2, 3, 4, 5, 6}
var i int
for i < len(numbers) {
    if i < len(numbers) {
        fmt.Println(numbers[i])
    }
    i++
}
```

Notice here that we are updating i in the for loop and omitted the init and post statements

We can even drop all statements and have an infinite loop

```go
var floats = []float64{1.5, 2.5, 3.5, 4.5, 5.5}
counter := 0

// Like While (true) Loop
for {
    if floats[counter] == 3.5 {
        break
    }
    counter++
}
```

Here it is our responsibility to break execution else run into a never terminating program

We can also use a `range` clause in the `for` loop

```go
var colors = []string{"red", "green", "blue", "yellow", "black"}
for idx, val := range colors {
    fmt.Println(idx, val)
}
```

The `range` clause returns 2 values in this loop:

1. `idx` is the index of the element of the slice
2. `val` is the current value

If you don't need one of the values returned you can use the `_` in Golang to simply ignore it

```go
var colors = []string{"red", "green", "blue", "yellow", "black"}
for _, val := range colors {
    fmt.Println(val)
}
```

Notice here we used the `_` and are ignoring the index returned to simply print the value at each index of the slice

[Loops Playground](https://play.golang.org/p/F6MC8Dktu1)