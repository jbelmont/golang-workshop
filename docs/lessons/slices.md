# Slices

[Slices Go Specification](https://golang.org/ref/spec#Slice_types)

Why Slices are preferred over Arrays in Golang?

* Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. 
* Except for items with explicit dimensions such as transformation matrices
* Most array programming in Go is done with slices rather than simple arrays.

More concrete Definition of Slices.

* Slices represent variable-length sequences whose elements all have the same type. 
* A slice type is written []T, where the elements have type T; it looks like an array type without a size.

* Arrays and slices are intimately connected. 
* A slice is a small data structure that gives access to a subsequence (or perhaps all) of the elements of an array, 
* This is known as the slice’s underlying array. 

A slice has three components: 

1. Pointer
2. Length 
3. Capacity

* A slice does not store any data, it just describes a section of an underlying array.
* Changing the elements of a slice modifies the corresponding elements of its underlying array.
* Other slices that share the same underlying array will see those changes.

```go
fruits := []string{
    "apples",
    "oranges",
    "strawberries",
    "tangerines",
    "cantaloupe",
}

fruits = append(fruits, "kiwi")
```

Notice in this example that we did not specify a length for the slice here.
Also notice that we left a trailing comma in the last value

If you didn't insert the last comma you would see the following error:

`syntax error: unexpected semicolon or newline, expecting comma or }`

The Go Specification says the following:

The formal grammar uses semicolons ";" as terminators in a number of productions. Go programs may omit most of these semicolons using the following two rules:

When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is:

1. an identifier
2. an integer, floating-point, imaginary, rune, or string literal
3. one of the keywords break, continue, fallthrough, or return
4. one of the operators and delimiters ++, --, ), ], or }

Notice here that we have a `string literal`

[Semicolons](https://golang.org/ref/spec#Semicolons)

```go
// Here numbers has length of 6 and capacity of 6
numbers := []int{2, 3, 5, 7, 11, 13}

// Slice the slice to give it zero length.
numbers = numbers[:0] // []

// Extend its length.
numbers = numbers[:4] // Sub index of [2 3 5 7]

// Drop its first two values.
numbers = numbers[2:] // [5 7 11 13]
```

* Notice here that we are using a colon operator in brackets, with it we can acccomplish sub-indexing
* Each time we use `[:]` we are reslicing the slice

```go
numbers = append(numbers, 55)
```

Notice here use the builtin `append` function this will dynamically increase the size of the slice by amount twice

If we were to print its length we would see 7 but a capacity of 12 since 6 * 2 = 12

We can also use the `make` builtin function for slices

```go
numbers = make([]int, 0, 15)
```

Here we created another numbers slice with a length of 0 but a capacity of 15 with the builtin `make` function

[Slices Playground](https://play.golang.org/p/6HNFnOiwJj)