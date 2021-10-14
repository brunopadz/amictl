package utils

import (
	"testing"
)

func TestEmptyString(t *testing.T) {
	a := ""

	if EmptyString(a) != "-" && len(a) == 0 {
		t.Errorf("EmptyString(\"\") = %v; want \"-\"", EmptyString(a))
	}

}
