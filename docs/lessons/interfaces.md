# Interfaces

[Interfaces Go Specification](https://golang.org/ref/spec#Interface_types)

* An interface type specifies a method set called its interface. 
* A variable of interface type can store a value of any type with a method set that is any superset of the interface. 
* Such a type is said to implement the interface. 
* The value of an uninitialized variable of interface type is nil.

* Interface types express abstractions about the behaviors of other types. 
* Interfaces let us write functions that are more flexible because they are not tied to the details of one particular implementation.

* Many object-oriented languages have some notion of interfaces
* Interfaces in Go are distinct because they are satisfied implicitly. 
* You don't have to declare all the interfaces that a given concrete type satisfies
* Possessing the necessary methods is enough

```go
type Shape interface {
	area()
}
```

Here we have declared an interface of `type` Shape that has methods area, circumference, and volume

Now we can define structs of type Circle and Rectangle

A Circle struct has the following fields:

```go
type Circle struct {
    pi float64
    radius float64
}
```

A Rectangle struct has the following fields:

```go
type Rectangle struct {
    height float64
    width float64
}
```

Now we can define an area method for the `circle` struct

```go
func (c Circle) area() float64 {
    return c.pi * math.Pow(c.radius, 2)
}
```

Now in order to declare a circle struct you can do this

```go
circle := Circle{
    radius: 5,
    pi: math.Pi
}
circle.area()
```

Notice here that we first created a circle struct and then called the `area` function on the circle struct

Let us define an area method on the Rectangle struct type

```go
func (r Rectangle) area() float64 {
    return r.height * r.width
}
```

Now that we have defined an `area` method on the Rectangle struct type let us finish

```go
rectangle := Rectangle{
    height: 4,
    width: 5,
}

rectangle.area()
```

[Interfaces Playground](https://play.golang.org/p/Emm21yKA48)