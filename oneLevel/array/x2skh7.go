package main

import "fmt"

/**
 * @note: 这个问题的关键在于 移动数据但是不能使得数据内存地址更改 想当于不能获得的结果不能变了数组
 * @auth: tongWz
 * @date: 2022年6月21日14:31:06
**/
func rotate(nums []int, k int) {
	//todo:1 这个方法可以 但是数组长的话就出问题
	//k = k % len(nums)
	//for ii := len(nums) - 1; ii >= k ; ii -- {
	//	nums[ii], nums[ii - k + 1] = nums[ii - k + 1], nums[ii]
	//}
	// todo：2 这个是可以的 内存地址不会被清除掉
	k = k % len(nums)
	tmp := append(nums[len(nums) - 1:], nums[:len(nums) - 1]...)
	nums = append(nums[:0], tmp...)
}

func main()  {
	var nums = []int{1,2,3,4,5,6,7}
	rotate(nums, 1)
	fmt.Println(nums)
}

//旋转数组
//给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//
// 
//
//示例 1:
//
//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右轮转 1 步: [7,1,2,3,4,5,6]
//向右轮转 2 步: [6,7,1,2,3,4,5]
//向右轮转 3 步: [5,6,7,1,2,3,4]
//
//作者：力扣 (LeetCode)
//链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2skh7/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
