# Arrays

[Arrays Go Specification](https://golang.org/ref/spec#Array_types)

* Arrays are useful when planning the detailed layout of memory and sometimes can help avoid allocation
* Arrays are the building block for slices.

There are major differences between the ways arrays work in Go and C languages
* Arrays are values. Assigning one array to another copies all the elements.
* If you pass an array to a function, it will receive a copy of the array, not a pointer to it.
* The size of an array is part of its type.
    * The types [10]int and [20]int are distinct

An array is a numbered sequence of elements of a single type with a fixed length.

```go
var numbers [5]int
numbers[0] = 1
numbers[1] = 2
numbers[2] = 3
numbers[3] = 4
numbers[4] = 5
```

Here is a declare an array of fixed length of 5 and individually assign `int` values to each index of the array.

If I were to assign another value `numbers[5] = 6` I would see the following error in Golang

`invalid array index 5 (out of bounds for 5-element array)`

[Arrays Playground](https://play.golang.org/p/a5nTP6ZI_R)
