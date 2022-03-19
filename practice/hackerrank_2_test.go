package practice

import (
	"testing"
)

func TestSimpleArraySumMain_success(t *testing.T) {
	input := "1,2,3,4,10,11"
	var expected int32 = 31
	result := SimpleArraySumMain(input)
	AssertEqual(expected, result, t)
}

func TestSimpleArraySumMain_inputContainLetter(t *testing.T) {
	input := "1,2,3,4,a,11"
	var expected int32 = 21
	result := SimpleArraySumMain(input)
	AssertEqual(expected, result, t)
}

func TestSimpleArraySumMain_inputIsNotNumeric(t *testing.T) {
	input := "a,b,c"
	var expected int32 = 0
	result := SimpleArraySumMain(input)
	AssertEqual(expected, result, t)
}
