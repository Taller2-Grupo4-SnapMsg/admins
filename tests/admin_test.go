package main

import "testing"

func TestDummy(t *testing.T) {
	// This is a dummy test that always passes
	if 1+1 != 2 {
		t.Errorf("1+1 should equal 2")
	}
}
