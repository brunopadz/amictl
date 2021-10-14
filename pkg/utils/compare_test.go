package utils

import (
	"testing"
)

func TestCompare(t *testing.T) {
	a := []string{"foo", "bar", "baz"}
	b := []string{"bar", "baz"}

	r := Compare(a, b)
	if r[0] != a[0] {
		t.Errorf("Compare(a, b) = %v,; want \"foo\"", Compare(a, b))
	}
}
