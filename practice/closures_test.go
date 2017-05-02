package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	numbers := []float64{
		1, 2, 3, 4, 5, 6, 7,
	}
	addEmUp := averageEmUp(numbers)
	if addEmUp != 4 {
		t.Error("Expected 4 but received ", addEmUp)
	}
}

func TestStandardDeviation(t *testing.T) {
	numbers := []float64{
		5.3,
		2.5,
		3.75,
		9.55,
		8.5,
		6.75,
	}
	actual := standardDeviation(numbers)()
	expected := "2.49"
	if actual != expected {
		t.Error("Expected 2.49 but received ", actual)
	}
}

func TestStandardDeviation2(t *testing.T) {
	numbers := []float64{
		5.3,
		2.5,
		3.75,
		9.55,
		8.5,
		6.75,
	}
	actual := standardDeviation(numbers)()
	expected := "2.49"
	assert.Equal(t, actual, expected)
}
