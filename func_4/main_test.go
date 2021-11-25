package func_4

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1, 1)
	want := 2

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
