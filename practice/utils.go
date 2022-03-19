package practice

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertToIntArray(arrInput string) []int32 {
	inputArrStr := strings.Split(arrInput, ",")
	inputArr := make([]int32, len(inputArrStr))
	for i, value := range inputArrStr {
		intValue, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			fmt.Printf("convert number failed for %v : %v\n", value, err)
		}

		inputArr[i] = int32(intValue)
	}
	return inputArr
}

func ConvertStringToNumber(input string) int32 {
	intValue, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		fmt.Printf("convert number failed for %v : %v\n", input, err)
	}
	return int32(intValue)
}

func ConvertToInt64Array(arrInput string) []int64 {
	inputArrStr := strings.Split(arrInput, ",")
	inputArr := make([]int64, len(inputArrStr))
	for i, value := range inputArrStr {
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			fmt.Printf("convert number failed for %v : %v\n", value, err)
		}

		inputArr[i] = int64(intValue)
	}
	return inputArr
}
