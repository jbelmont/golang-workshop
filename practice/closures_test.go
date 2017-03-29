package practice

import (
	"testing"
)

func TestAverage(t *testing.T) {
	numbers := []int{
		1, 2, 3, 4, 5, 6, 7,
	}
	addEmUp := averageEmUp(numbers)
	if addEmUp != 4 {
		t.Error("Expected 4 but received ", addEmUp)
	}
}
