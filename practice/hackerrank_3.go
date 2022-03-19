package practice

func CompareTriplets(a string, b string) []int32 {
	var result []int32 = make([]int32, 2)

	calculatePoint(convertToIntArray(a), convertToIntArray(b), 0, result)

	return result
}

func calculatePoint(a []int32, b []int32, i int, result []int32) {
	pointA := 0
	pointB := 0
	if i == len(a) {
		return
	}

	if a[i] > b[i] {
		pointA += 1
	} else if a[i] < b[i] {
		pointB += 1
	}
	result[0] = result[0] + int32(pointA)
	result[1] = result[1] + int32(pointB)
	calculatePoint(a, b, i+1, result)
}
