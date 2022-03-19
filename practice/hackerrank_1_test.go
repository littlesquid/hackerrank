package practice

import (
	"testing"
)

func TestSolveMeFirst(t *testing.T) {
	var expected uint32 = 5
	result := SolveMeFirst(2, 3)
	AssertEqual(expected, result, t)
}
