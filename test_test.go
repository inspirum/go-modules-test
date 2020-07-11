package test

import "testing"

func TestFunc(t *testing.T) {
	got := Test1(1, 3)
	if got != 4 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
