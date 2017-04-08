# Files

* The io package consists of a few functions, but mostly interfaces used in other packages. 
* The two main interfaces are Reader and Writer. 
* Readers support reading via the Read method. 
* Writers support writing via the Write method. 
* Many functions in Go take Readers or Writers as arguments. 

* To open a file in Go use the Open function from the os package.

[OS Go Package](https://golang.org/pkg/os/)

[IOUTIL Go Package](https://golang.org/pkg/io/ioutil/)

```go
soldiersTxt, err := os.Create("soldiers.txt")
if err != nil {
    log.Fatal(err)
}
```

Here we have a call to `os.Create` which creates a file

```go
lines := []byte("I am a string\nWhat is the string\nNever is the string")
err8 := ioutil.WriteFile("soldiers.txt", lines, 0644)
if err8 != nil {
    log.Fatal(err8)
}
assert(err8 != nil)

fileJSON, err9 := ioutil.ReadFile("data.json")
if err9 != nil {
    log.Fatal(err9)
}
```

Please read up on the `os` and `ioutil` packages for more information on files