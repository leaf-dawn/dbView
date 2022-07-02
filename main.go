package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(9))

}

//二分查找
//建立一个表,用于判断是否是
//二分一波
func isPerfectSquare(num int) bool {
	left := 1
	right := num

	for left < right {
		mid := (left + right) / 2

		if mid*mid > num {
			right = mid - 1
		} else {
			left = left + 1
		}
	}

	return left*left == num

}
