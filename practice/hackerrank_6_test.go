package practice

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPlusMinus_case1(t *testing.T) {
	input := []int32{1, 1, 0, -1, -1}

	positiveRatioExpect := "0.400000"
	negativeRatioExpect := "0.400000"
	zeroRatioExpect := "0.200000"

	result := PlusMinus(input)
	positiveRatio := result[0]
	negativeRatio := result[1]
	zeroRatio := result[2]

	AssertEqual(positiveRatioExpect, positiveRatio, t)
	AssertEqual(negativeRatioExpect, negativeRatio, t)
	AssertEqual(zeroRatioExpect, zeroRatio, t)
}

func TestPlusMinus_case2(t *testing.T) {
	input := []int32{-4, 3, -9, 0, 4, 1}

	positiveRatioExpect := "0.500000"
	negativeRatioExpect := "0.333333"
	zeroRatioExpect := "0.166667"

	result := PlusMinus(input)
	positiveRatio := result[0]
	negativeRatio := result[1]
	zeroRatio := result[2]

	AssertEqual(positiveRatioExpect, positiveRatio, t)
	AssertEqual(negativeRatioExpect, negativeRatio, t)
	AssertEqual(zeroRatioExpect, zeroRatio, t)
}

func FuzzPlusMinus(f *testing.F) {
	testcases := [][]int32{
		{-4, 3, -9, 0, 4, 1},
		{1, 1, 0, -1, -1},
		{-9, -1, 0, 12, 8, 1, 0},
	}

	for _, testcase := range testcases {
		data, _ := json.Marshal(testcase)
		f.Add(data)
	}

	f.Fuzz(func(t *testing.T, a []byte) {
		var input []int32

		err := json.Unmarshal(a, &input)
		if err != nil {
			fmt.Println("parse input failed")
		}
		PlusMinus(input)
	})
}
