package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	out1 := isRobotBounded("GGLLGG")
	want := true
	if out1 != want {
		t.Errorf("got %t, want %t", out1, want)
	}
}
