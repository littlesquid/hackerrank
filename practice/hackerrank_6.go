package practice

import (
	"fmt"
)

func PlusMinus(arr []int32) []string {
	ratioList := checkRatio(arr, 0, [3]float32{})
	size := float32(len(arr))
	positiveRatio := fmt.Sprintf("%.6f", ratioList[0]/size)
	negativeRatio := fmt.Sprintf("%.6f", ratioList[1]/size)
	zeroRatio := fmt.Sprintf("%.6f", ratioList[2]/size)

	result := []string{positiveRatio, negativeRatio, zeroRatio}

	return result
}

func checkRatio(arr []int32, index int, ratioResult [3]float32) [3]float32 {

	if index == len(arr) {
		return ratioResult
	}

	if arr[index] > 0 {
		ratioResult[0] += 1
	} else if arr[index] < 0 {
		ratioResult[1] += 1
	} else {
		ratioResult[2] += 1
	}

	return checkRatio(arr, index+1, ratioResult)
}
