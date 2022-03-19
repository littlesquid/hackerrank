package hackerrank

import (
	"testing"
)

func AssertEqual(expected interface{}, actual interface{}, t *testing.T) {
	if actual != expected {
		t.Errorf("result should be %v, actual result is %v", expected, actual)
	}
}
