package template

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	name := templates()
	if name != "John Rambo" {
		t.Errorf("name should equal John Rambo")
	}
}
