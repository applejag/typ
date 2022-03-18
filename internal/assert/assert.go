package assert

import "testing"

func Comparable[T comparable](t *testing.T, name string, want T, got T) {
	if want != got {
		t.Errorf(`%s: want "%v", got "%v"`, name, want, got)
	}
}
