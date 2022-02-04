package main

import (
	// "fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	v := Average(1,  2)

	if v != 1.5 {
		t.Errorf("expected 1.5 got %v", v)
	}
}
