package practice

import "testing"

func TestVeryBigSum(t *testing.T) {
	input := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	var expected int64 = 5000000015
	actual := VeryBigSum(input)
	AssertEqual(expected, actual, t)
}

func TestVeryBigSum_case2(t *testing.T) {
	input := []int64{1000000003234531, 13344000000002, 1003420000003, 10035440000004, 10000021220005}
	var expected int64 = 1034382884454545
	actual := VeryBigSum(input)
	AssertEqual(expected, actual, t)
}
