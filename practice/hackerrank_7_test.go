package practice

import "testing"

func FuzzStaircase(f *testing.F) {
	f.Add(int32(10))
	f.Fuzz(func(t *testing.T, n int32) {
		Staircase(n)
	})
}
