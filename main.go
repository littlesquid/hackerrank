package main

import (
	"fmt"
	"hackerrank/practice"
	"log"
	"strings"
)

func main() {
	var practiceNo int8
	titles := []string{
		"1. SolveMeFirst",
		"2. SimpleArraySum",
		"3. CompareTriplets",
		"4. VeryBigSum",
		"5. DiagonalDifference",
	}
	for _, title := range titles {
		fmt.Println(title)
	}
	fmt.Print("Enter practice number: ")
	fmt.Scanf("%v", &practiceNo)

	switch practiceNo {
	case 1:
		prepareHackerRank1()
	case 2:
		prepareHackerRank2()
	case 3:
		prepareHackerRank3()
	case 4:
		prepareHackerRank4()
	case 5:
		prepareHackerRank5()
	}

}

func prepareHackerRank1() {
	//uint = unsigned integer can hold zero and positive number
	var a, b, res uint32
	fmt.Print("Enter the number:")
	fmt.Scanf("%v+%v", &a, &b)
	res = practice.SolveMeFirst(a, b)

	fmt.Printf("result is %v", res)
}

func prepareHackerRank2() {
	var input string
	fmt.Print("Enter the numbers:")
	fmt.Scanf("%v", &input)
	practice.SimpleArraySumMain(input)
}

func prepareHackerRank3() {
	var aliceInput string
	var bobInput string
	fmt.Print("Alice - Enter the numbers:")
	fmt.Scanf("%v", &aliceInput)
	fmt.Print("Bob - Enter the numbers:")
	fmt.Scanf("%v", &bobInput)

	result := practice.CompareTriplets(aliceInput, bobInput)

	fmt.Printf("result is %v", result)
}

func prepareHackerRank4() {
	var input string
	fmt.Print("Enter the large numbers:")
	fmt.Scanf("%v", &input)
	var inputArr []int64 = practice.ConvertToInt64Array(input)
	result := practice.VeryBigSum(inputArr)

	fmt.Printf("result is %v", result)
}

func prepareHackerRank5() {
	var size int
	fmt.Print("Enter the matrix size (odd number): ")
	fmt.Scanf("%v", &size)

	if size%2 != 0 {
		log.Fatal("metrix size should be odd number")
	}

	r := 0
	var userInput string
	inputs := [][]int32{}
	for r < size {
		fmt.Printf("Enter the set or number for row%v: ", r)
		fmt.Scanf("%v", &userInput)
		c := 0
		inputArr := strings.Split(userInput, ",")
		for c < len(inputArr) {
			inputs[r][c] = practice.ConvertStringToNumber(inputArr[c])
			c++
		}
		r++
	}

	result := practice.DiagonalDifference(inputs, size)
	fmt.Printf("result is %v", result)
}
