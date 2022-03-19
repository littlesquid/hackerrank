package hackerrank

import (
	"fmt"
	"strconv"
	"strings"
)

func simpleArraySum(arr []int32) int32 {
	var result int32 = 0
	for _, value := range arr {
		result += value
	}
	return result
}

func SimpleArraySumMain(input string) int32 {
	inputArrStr := strings.Split(input, ",")
	inputArr := make([]int32, len(inputArrStr))
	for i, value := range inputArrStr {
		intValue, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			fmt.Printf("convert number failed for %v\n", value)
		}

		inputArr[i] = int32(intValue)
	}

	result := simpleArraySum(inputArr)

	fmt.Printf("summary of [%v] is %v\n", input, result)
	return result
}
