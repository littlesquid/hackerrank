package practice

import (
	"testing"
)

func TestSolveMeFirst(t *testing.T) {
	var expected uint32 = 5
	result := SolveMeFirst(2, 3)
	AssertEqual(expected, result, t)
}

func FuzzSolveMeFirst(f *testing.F) {
	f.Add(uint32(5), uint32(7))
	f.Fuzz(func(t *testing.T, a uint32, b uint32) {
		SolveMeFirst(a, b)
	})
}
