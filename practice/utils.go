package practice

import (
	"fmt"
	"strconv"
	"strings"
)

func convertToIntArray(arrInput string) []int32 {
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
