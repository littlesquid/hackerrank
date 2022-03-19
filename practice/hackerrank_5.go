package practice

import (
	"math"
)

/*
input:
1. n = number of row and column of metrix
2. loop on n and input n number separate by, for each row
*/

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
