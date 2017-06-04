# Variadic Functions

* A variadic function is one that can be called with varying numbers of arguments. 
* The most familiar examples are fmt.Printf and its variants. 
* Printf requires one fixed argument at the beginning, then accepts any number of subsequent arguments
* To create a variadic function, the type of the final parameter is preceded by an ellipsis, "..." 
    * This indicates that the function may be called with any number of arguments of this type.

```go
func sum(vals ...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}

values := []int{1, 2, 3, 4}
fmt.Println(sum(values...)) // "10"
```

Notice here the `...` after the parameter vals and notice the `...` after values in the caller
Think of this like a spread operator which spreads all the arguments of a list

[Variadic Functions Playground](https://play.golang.org/p/6pAWZDyzLM)