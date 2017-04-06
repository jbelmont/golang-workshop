# Pointers

[Pointers Go Specification](https://golang.org/ref/spec#Pointer_types)

* Go has pointers. A pointer holds the memory address of a value.
* The type *T is a pointer to a T value. Its zero value is nil.
* A pointer type denotes the set of all pointers to variables of a given type, called the base type of the pointer.
* The value of an uninitialized pointer is nil.
* In Go you cannot do pointer arithmetic however

`var p *int` returns `nil`

* The & operator generates a pointer to its operand.
* The * operator denotes the pointer's underlying value.

```go
var num = 3
otherNum := &num // otherNum now points to memory address of num
*otherNum++

fmt.Println(*otherNum, num) // now contain the value 4
```

Note the use of the `*otherNum` operator this is known as deferencing while the `&num` points to memory address

```go
type Number struct {
	num int
}

func (n *Number) scaler(s int) int {
	n.num *= s
	return n.num
}
``` 

[Pointer Playground](https://play.golang.org/p/X85mr4IbMm)