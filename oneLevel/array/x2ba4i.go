package main

import "fmt"

func moveZeroes(nums []int) {
	zeroT := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && i < len(nums)-1 && len(nums)-i > zeroT {
			nums = append(append(nums[:i], nums[(i+1):]...), 0)
			// 如果这个值发现是0 那么我们还在同样的位置还要看提前过来的是不是0
			i--
			// 如果最后几个一直是0我们不能无限循环
			zeroT++
		}
	}
}

func main() {
	var tSlice = []int{0, 1, 0, 3, 12}
	var tSlice2 = []int{0, 0, 1}
	moveZeroes(tSlice)
	moveZeroes(tSlice2)
	fmt.Printf("值是：%#v, \n %#v \n", tSlice, tSlice2)
}

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//请注意，必须在不复制数组的情况下原地对数组进行操作。
//
//
//
//示例 1:
//
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
//示例 2:
//
//输入: nums = [0]
//输出: [0]
//
//作者：力扣 (LeetCode)
//链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2ba4i/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
