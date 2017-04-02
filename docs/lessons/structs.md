# Structs

A struct is a collection of fields.

Struct can have methods to them which are similar to functions

* A struct is a sequence of named elements, called fields each of which has a name and a type. 
* Field names may be specified explicitly (IdentifierList) or implicitly (AnonymousField).
* Within a struct, non-blank field names must be unique.

[Structs Go Specification](https://golang.org/ref/spec#Struct_types)

Here is an example of an empty struct in Golang

```go
type Empty struct{}
var e Empty
fmt.Println(e) // {}
```

Here we create an empty struct and then declare a variable named `e` of type `empty`

```go
type Student struct {
    name  string
	level string
    gpa string
}
```

Here is a `Student` struct with fields `name`, `level`, and `gpa`

```go
student := Student{
    "Jane Doe",
    "Sophomore",
    "4.0",
}
```

Remember that Go will complain about the trailing comma for string literals hence why we added it here

`fmt.Println(student)` would show the contents of `{Jane Doe Sophomore 4}`

Struct fields are accessed using a dot.

```go
fmt.Println(student.level) // Sophomore
```

### Methods

* A method is declared with a variant of the ordinary function declaration
* An extra parameter appears before the function name. 
* The parameter attaches the function to the type of that parameter.

```go
type Student struct {
    Name  string
	Level string
    Gpa string
}

func (s Student) PrintReport() {
    fmt.Printf("%s is a %s student who is a %s", s.Name, s.Gpa, s.Level)
}
```

`PrintReport` looks like an ordinary function but has an extra parameter that appears before the function name.

* The extra parameter p is called the method’s receiver
* Dates back from early object-oriented languages that described calling a method as "sending a message to an object."
    * Smalltalk being one of them

* In Go, we don’t use a special name like `this` or `self` for the receiver 
* Here we choose receiver names just as we would for any other parameter. 
* Since the receiver name will used often, choose something short and be consistent across methods.
* A common choice is the first letter of the type name, like s for Student.

* In a method call, the receiver argument appears before the method name. 
* This parallels the declaration, in which the receiver parameter appears before the method name.

Here is how we would call the PrintReport method

```go
student := Student{
    "Jane Doe",
    "Sophomore",
    "4.0",
}

student.PrintReport() // Jane Doe is a Sophomore who is a 4.0 student
```

[Structs Playground](https://play.golang.org/p/1VJefxgw6E)