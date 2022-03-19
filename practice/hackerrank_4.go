package practice

func VeryBigSum(arr []int64) int64 {
	sum := int64(0)
	for _, v := range arr {
		sum += v
	}
	return sum
}
