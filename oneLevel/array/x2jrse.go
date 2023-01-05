package main

import "fmt"

func twoSum(nums []int, target int) []int {
	needV := make(map[int]int)
	for k, v := range nums {
		if v <= target {
			needV[target-v] = k
		}
	}
	for k, v := range nums {
		if index, ok := needV[v]; ok && k != index {
			return []int{k, index}
		}
	}
	return nil
}

func main() {
	var sliceT = []int{2, 7, 11, 15}
	fmt.Printf("值是：%#v", twoSum(sliceT, 9))
}
