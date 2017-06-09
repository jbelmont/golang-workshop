package summer

import (
	"testing"
)

func TestSummer(t *testing.T) {
	numbers := []float64{
		1.5, 2.5, 3.75, 2.8, 9.8, 1111.5,
	}
	actual := Sum(numbers)
	expected := 1131.85
	if actual != expected {
		t.Errorf("should be equal to %f", expected)
	}
}
