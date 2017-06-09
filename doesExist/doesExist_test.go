package doesExist

import (
	"testing"
)

func TestDoesExist(t *testing.T) {
	fruits := []string{
		"apple", "orange", "kiwi", "pineapple",
	}
	const mango = "mango"
	if Exists(fruits, mango) {
		t.Errorf("%s should not exist", mango)
	}
}
