package workshop

import "fmt"

func builtins() {
	var slice1 = []string{
		"I am a string",
		"I am the other string",
	}
	slice := append(slice1, "Just another string")
	sliceFinal := make([]string, len(slice))
	copy(sliceFinal, slice)
	comparator := func(slice, slicefinal []string) bool {
		fmt.Println(slice, sliceFinal)
		for i := range slice {
			if slice[i] != slicefinal[i] {
				return false
			}
		}
		return true
	}
	assert(comparator(slice, sliceFinal) == true)
}
