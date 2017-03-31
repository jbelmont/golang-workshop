package workshop

import (
	"fmt"
	"testing"
)

func TestConditionals(t *testing.T) {
	variables()
	conditionals()
	loops()
	functions()
	arrays()
	slices()
	pointers()
	maps()
	structs()
	interfaces()
	builtins()
	jsonExamples()
	files()

	fmt.Printf("\n%c[32;1mCongratulations you completed the Workshop!!!%c[0m\n\n", 27, 27)
}
