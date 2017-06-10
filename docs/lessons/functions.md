# Functions

[Functions Go Specification](https://golang.org/ref/spec#Function_types)

* A function lets us wrap up a sequence of statements as a unit that can be called from elsewhere in a program
    * Sometimes a function is used repeatedly but not always

Example function declaration

```go
func name(parameter-list) (result-list) {
    body
}
```

* declare a function in Go by using the keyword `func`
* name the function using legal names
    * The spec says that func, var or const name must begin with (unicode_letter or _)
    * and can end with many (unicode_letter, unicode_digit or _)
    * The blank identifier, represented by the underscore character _, may be used in a declaration like any other identifier but the declaration does not introduce a new binding.
* The parameter list specifies the names and types of the function’s parameters,
    * These are local variables whose values or arguments are supplied by the caller of the function
* If the function returns one unnamed result or no results at all, parentheses are optional and usually omitted

## Multiple Return Values

Functions in Golang can return multiple values

```go
func giveErrorIfEmpty(str string) (string, error) {
	if str == "" {
		error := errors.New("You must pass a string")
		return str, error
	}
	return str, nil
}
```

This function returns either the passed `str` parameter and `nil` error or `str` with an `error`

## Variadic Functions

A variadic function is one that can be called with varying numbers of arguments.

* To declare a variadic function, the type of the final parameter is preceded by an ellipsis, `...`
* This indicates that the function may be called with any number of arguments of the provided type

```go
func summation(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}
```

`summation` can take 0 or more arguments of integer values and will sum up those values

## Anonymous Functions

* A function literal is written like a function declaration, but without a name following the func keyword.
* A function literal is an expression, and its value is called an anonymous function.
* Functions defined in this way have access to the entire lexical environment,
* so the inner function can refer to variables from the enclosing function
    * Otherwise known as a closure

```go
func distanceOfLine(x1, y1, x2, y2 float64) func() string {
	var distance float64
	return func() string {
		distance = math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
		return format2DecimalPlaces(distance)
	}
}
```

The function `distanceOfLine` has an anonymous function inside of it
The variable distance is available to the inner function and one of the advantages of closures is they retain `state`

## Recursion

Functions may be recursive, that is, they may call themselves, either directly or indirectly.

```go
func summationRecursion(number int) int {
	if number > 0 {
		return number + summationRecursion(number-1)
	}
	return number
}
```

Here is an example of recursion in Golang

*Remember with recursion there must exist a base case condition to prevent an infinite call stack*

[Functions Playground](https://play.golang.org/p/twEJHvYBlz)
