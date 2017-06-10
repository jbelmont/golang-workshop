package reverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	actual := reverse("gorilla")
	expected := "allirog"
	if actual != expected {
		t.Errorf("%s should equal %s", actual, expected)
	}
}
