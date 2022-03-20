package practice

import (
	"math"
)

/*
input:
1. n = number of row and column of metrix
2. loop on n and input n number separate by, for each row
*/

func diagonalSummary(inputs [][]int32, row, col int, sum int32, diagonalType string) int32 {

	size := int(len(inputs))

	if row == size && col == 0 {
		return sum
	}

	if col >= 0 {
		switch diagonalType {
		case "primary":
			if row == col {
				sum += inputs[row][col]
			}
		case "secondary":
			if (row + col) == (size - 1) {
				sum += inputs[row][col]
			}
		}
		return diagonalSummary(inputs, row, col-1, sum, diagonalType)
	} else {
		return diagonalSummary(inputs, row+1, size-1, sum, diagonalType)
	}
}

func PrimaryDiagonal(inputs [][]int32, size int) int32 {
	sum := int32(0)
	for r := 0; r < size; r++ {
		for c := size - 1; c >= 0; c-- {
			if r == c {
				sum += inputs[r][c]
			}
		}
	}
	return sum
}

func SecondaryDiagonal(inputs [][]int32, size int) int32 {
	sum := int32(0)
	for r := 0; r < size; r++ {
		for c := size; c >= 0; c-- {
			if r+c == size-1 {
				sum += inputs[r][c]
			}
		}
	}
	return sum
}

func DiagonalDifference(inputs [][]int32) int32 {
	size := len(inputs)
	primary := float64(PrimaryDiagonal(inputs, size))
	secondary := float64(SecondaryDiagonal(inputs, size))
	return int32(math.Abs(primary - secondary))
}

func DiagonalDifferenceRecursion(inputs [][]int32) int32 {
	size := len(inputs)
	sum := int32(0)
	primary := float64(diagonalSummary(inputs, 0, size-1, sum, "primary"))
	secondary := float64(diagonalSummary(inputs, 0, size, sum, "secondary"))
	return int32(math.Abs(primary - secondary))
}
