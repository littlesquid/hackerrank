package practice

import (
	"testing"
)

func TestDiagonalDifference3x3(t *testing.T) {
	input := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	expected := int32(15)

	actualResult := DiagonalDifference(input)

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

	actualResult := DiagonalDifference(input)

	AssertEqual(expected, actualResult, t)
}

func TestDiagonalDifference4x4_1(t *testing.T) {
	input := [][]int32{
		{-1, 1, -7, -8},
		{-10, -8, -5, -2},
		{0, 9, 7, -1},
		{4, 4, -2, 1},
	}

	expected := int32(1)

	actualResult := DiagonalDifference(input)

	AssertEqual(expected, actualResult, t)
}

func TestDiagonalDifference9x9(t *testing.T) {
	input := [][]int32{
		{6, 6, 7, -10, 9, -3, 8, 9, -1},
		{9, 7, -10, 6, 4, 1, 6, 1, 1},
		{-1, -2, 4, -6, 1, -4, -6, 3, 9},
		{-8, 7, 6, -1, -6, -6, 6, -7, 2},
		{-10, -4, 9, 1, -7, 8, -5, 3, -5},
		{-8, -3, -4, 2, -3, 7, -5, 1, -5},
		{-2, -7, -4, 8, 3, -1, 8, 2, 3},
		{-3, 4, 6, -7, -7, -8, -3, 9, -6},
		{-2, 0, 5, 4, 4, 4, -3, 3, 0},
	}

	expected := int32(52)

	actualResult := DiagonalDifference(input)

	AssertEqual(expected, actualResult, t)
}
