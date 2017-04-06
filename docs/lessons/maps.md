# Maps

[Make Go Specification](https://golang.org/ref/spec#Map_types)

* In Go, a map is a reference to a hash table, and a map type is written map[K]V
    * K and V are the types of its keys and values. 
* All of the keys in a given map are of the same type, and all of the values are of the same type, 
    * The keys need not be of the same type as the values. 
* The key type K must be comparable using ==, so that the map can test whether a given key is equal to one already within it.

* A map is an unordered group of elements of one type, called the element type
* indexed by a set of unique keys of another type, called the key type.

In other words a `map` maps keys to values.

* The zero value of a map is nil. A nil map has no keys, nor can keys be added.
* The make function returns a map of the given type, initialized and ready for use.

```go
type Students struct {
    ages map[string]int
}

type Coordinates struct {
    X,Y float64
}

ages := make(map[string]int, 5)
students := Students{ages}
students.ages["jane"] = 18
students.ages["jake"] = 21
```

Notice the `Students` struct has a field of `ages` which is a map type
We can use the `make` builtin function to create a new empty map value which takes the map type and a capacity

```go
type Coordinates struct {
    X,Y float64
}

var points = map[string]Coordinates{
    "Points 1": Coordinates{
        X: 2.5, Y: 3.75,
    },
}

var points2 = make(map[string]Coordinates)
points2["Points 2"] = Coordinates{
    X: 3.75, Y: 5.5,
}
```

* Notice the difference between `points2` and `points`. For `points` we are using a map literal
* Map literals are like struct literals, but the keys are required.
* For `points2` we used the make builtin and assigned a key value of `Points 2`

[Maps Playground](https://play.golang.org/p/lm8SxAKWez)