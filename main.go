package main

import (
	"fmt"
	"hackerrank/practice"
)

func main() {
	var practiceNo int8
	fmt.Print("1. SolveMeFirst\n2. SimpleArraySum\n3. CompareTriplets\n")
	fmt.Println("4. VeryBigSum")
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
