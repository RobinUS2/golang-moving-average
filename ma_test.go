package movingaverage

// @author Robin Verlangen
// Test the moving average for Go package

import (
	"testing"
)

func TestMovingAverage(t *testing.T) {
	a := New(5)
	if a.Avg() != 0 {
		t.Fail()
	}

	if a.SlotsFilled() {
		t.Error("should not be full yet")
	}

	a.Add(2)
	if a.Avg() < 1.999 || a.Avg() > 2.001 {
		t.Fail()
	}
	a.Add(4)
	a.Add(2)
	if a.Avg() < 2.665 || a.Avg() > 2.667 {
		t.Fail()
	}
	a.Add(4)
	a.Add(2)
	if a.Avg() < 2.799 || a.Avg() > 2.801 {
		t.Fail()
	}

	if !a.SlotsFilled() {
		t.Error("should be full")
	}

	// This one will go into the first slot again
	// evicting the first value
	a.Add(10)
	if a.Avg() < 4.399 || a.Avg() > 4.401 {
		t.Fail()
	}

	// test variadic add
	a.Add(2, 4)

	// get values
	values := a.Values()
	if len(values) != 5 {
		t.Error()
	}
}
