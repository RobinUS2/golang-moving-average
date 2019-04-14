package movingaverage

// @author Robin Verlangen
// Test the moving average for Go package

import (
	"math"
	"sync"
	"testing"
)

func TestMovingAverage(t *testing.T) {
	a := New(5)
	if a.Avg() != 0 {
		t.Error("expected 0", a.Avg())
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

func TestNaN(t *testing.T) {
	a := New(5)
	a.Add(1)
	a.Add(math.NaN())
	if !math.IsNaN(a.Avg()) {
		t.Error()
	}
}

func TestNaNIgnore(t *testing.T) {
	a := New(5)
	a.SetIgnoreNanValues(true)
	a.Add(1)
	a.Add(math.NaN())
	if math.IsNaN(a.Avg()) {
		t.Error()
	}
	if a.Avg() != 1 {
		t.Error(a.Avg())
	}
}

func TestInf(t *testing.T) {
	a := New(5)
	a.Add(1)
	a.Add(math.Inf(1))
	if !math.IsInf(a.Avg(), 0) {
		t.Error()
	}
}

func TestInfIgnore(t *testing.T) {
	a := New(5)
	a.SetIgnoreInfValues(true)
	a.Add(1)
	a.Add(math.Inf(1))
	if math.IsInf(a.Avg(), 0) {
		t.Error()
	}
	if a.Avg() != 1 {
		t.Error(a.Avg())
	}
}

func TestMax(t *testing.T) {
	a := New(5)
	if max, err := a.Max(); max != 0 || err == nil {
		t.Error(max, err)
	}
	a.Add(10)
	if max, err := a.Max(); max != 10 || err != nil {
		t.Error(max, err)
	}
	a.Add(100, 10000, 1000)
	if max, err := a.Max(); max != 10000 || err != nil {
		t.Error(max, err)
	}
}

func TestMin(t *testing.T) {
	a := New(5)
	if min, err := a.Min(); min != 0 || err == nil {
		t.Error(min, err)
	}
	a.Add(10)
	if min, err := a.Min(); min != 10 || err != nil {
		t.Error(min, err)
	}
	a.Add(100, 10000, 1000)
	if min, err := a.Min(); min != 10 || err != nil {
		t.Error(min, err)
	}
}

func TestCount(t *testing.T) {
	a := New(5)
	if a.Count() != 0 {
		t.Error(a.Count())
	}
	a.Add(5)
	if a.Count() != 1 {
		t.Error(a.Count())
	}
	a.Add(3, 6)
	if a.Count() != 3 {
		t.Error(a.Count())
	}
	a.Add(1, 2, 3, 4, 5)
	if a.Count() != 5 {
		t.Error(a.Count())
	}
}

func TestConcurrent(t *testing.T) {
	// this test needs to be run with -race flag
	a := Concurrent(New(5))

	const numRoutines = 5
	wg := sync.WaitGroup{}
	wg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func() {
			for n := 0; n < 10; n++ {
				a.Add(float64(n))
			}
			a.Avg()
			a.Min()
			a.Max()
			a.Count()
			a.Values()
			a.SlotsFilled()
			wg.Done()
		}()
	}
	wg.Wait()
}
