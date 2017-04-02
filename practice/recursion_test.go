package practice

import (
	"testing"
)

func TestSummationRecursion(t *testing.T) {
	value := summationRecursion(3)
	if value != 6 {
		t.Error("Expected 6 but received ", value)
	}
}
