package practice

import "fmt"

func Staircase(n int32) {
	value := ""
	print(n, 0, 0, value)
}

func print(n int32, row, col int32, value string) {
	if row == n && col == n {
		return
	}

	if row+col == n {
		value += "#"
	}

	if col < n {
		print(n, row, col+1, value)
	} else if col == n {
		fmt.Println(fmt.Sprintf("%*v", col, value))
		print(n, row+1, 0, value)
	}
}
