package variables

import (
	"fmt"
	"testing"
)

func TestVariables(t *testing.T) {
	v1, v2 := variables()
	fmt.Println(v1, v2)
}
