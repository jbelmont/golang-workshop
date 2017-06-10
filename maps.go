package workshop

import (
	"strconv"
)

func maps() {
	// declare and initialize map in one line
	var intmap = map[string]int{"one": 1, "two": 2, "three": 3}
	var one = intmap["one"]
	assert(one == 2)

	// Use the make function to create an empty map
	var strmap = make(map[string]string)
	strVal := "val"
	for i := 1; i <= 10; i++ {
		strmap[strVal+strconv.Itoa(i)] = strVal + " is the number " + strconv.Itoa(i)
	}
	assert(strmap["val3"] == "what am i")

	// delete key/value pair from map
	delete(intmap, "three")

	// len builtin on map reports the number of key/value pairs on map
	assert(len(intmap) == 3)
}
