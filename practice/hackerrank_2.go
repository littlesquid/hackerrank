package practice

import (
	"fmt"
)

func simpleArraySum(arr []int32) int32 {
	var result int32 = 0
	for _, value := range arr {
		result += value
	}
	return result
}

func SimpleArraySumMain(input string) int32 {
	inputArr := convertToIntArray(input)
	result := simpleArraySum(inputArr)
	fmt.Printf("summary of [%v] is %v\n", input, result)
	return result
}
