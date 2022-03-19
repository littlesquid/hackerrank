package main

import "fmt"

func main() {
	var practiceNo int8
	fmt.Print("1. SolveMeFirst\n2. SimpleArraySum\nEnter practice number: ")
	fmt.Scanf("%v", &practiceNo)

	switch practiceNo {
	case 1:
		prepareHackerRank1()
	case 2:
		prepareHackerRank2()
	}

}

func prepareHackerRank1() {
	//uint = unsigned integer can hold zero and positive number
	var a, b, res uint32
	fmt.Print("Enter the number:")
	fmt.Scanf("%v+%v", &a, &b)
	res = SolveMeFirst(a, b)

	fmt.Printf("result is %v", res)
}

func prepareHackerRank2() {
	var input string
	fmt.Print("Enter the numbers:")
	fmt.Scanf("%v", &input)
	SimpleArraySumMain(input)
}
