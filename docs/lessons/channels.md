# Channels

[Channels Go Specification](https://golang.org/ref/spec#Channel_types)

## About Channel Types

* Channels provide a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type. 
* The value of an uninitialized channel is nil.

* The optional <- operator specifies the channel direction, send or receive. 
* If no direction is given, the channel is bidirectional. 
* A channel may be constrained only to send or only to receive by conversion or assignment.

* The <- operator associates with the leftmost chan possible:

* A new, initialized channel value can be made using the built-in function make
    * which takes the channel type and an optional capacity as arguments:

```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```

* In this example we make a channel with the statement `c := make(chan int)`
* Notice the difference between the use `c <- sum` and `<-c`
    * `c <- sum` is sending values whereas `<c` is receiving values

[Channels Playground](https://play.golang.org/p/ORc2utkZAE)