package practice

import "testing"

func TestDiagonalDifference3x3(t *testing.T) {
	input := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	expected := int32(15)

	actualResult := DiagonalDifference(input, 3)

	AssertEqual(expected, actualResult, t)
}

func TestDiagonalDifference4x4(t *testing.T) {
	input := [][]int32{
		{11, 2, 4, 1},
		{4, 5, 6, 8},
		{10, 8, -12, -9},
		{15, 7, 2, 5},
	}

	expected := int32(21)

	actualResult := DiagonalDifference(input, 4)

	AssertEqual(expected, actualResult, t)
}
