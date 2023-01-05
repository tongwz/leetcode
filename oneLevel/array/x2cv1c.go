package main

import "fmt"

func plusOne(digits []int) []int {
	endIndex := len(digits) - 1
	if digits[endIndex] < 9 {
		digits[endIndex]++
	} else {
		if endIndex == 0 {
			digits = []int{1, 0}
		} else {
			digits = append(plusOne(digits[:endIndex]), 0)
		}

	}
	return digits
}

func main() {
	var intFlag = []int{1, 8, 9}
	fmt.Printf("值是：%#v", plusOne(intFlag))
}
