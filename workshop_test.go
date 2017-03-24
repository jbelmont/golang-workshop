package workshop

import (
	"fmt"
	"testing"
)

func TestConditionals(t *testing.T) {
	variables()
	conditionals()
	loops()
	structs()
	interfaces()

	fmt.Printf("\n%c[32;1mCongratulations you completed the Workshop!!!%c[0m\n\n", 27, 27)
}
