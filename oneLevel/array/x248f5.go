package main

func containsDuplicate(nums []int) bool {
	// todo:1 通过循环查询
	//len := len(nums)
	//for i := 0; i < len - 1; i++ {
	//	for j := i + 1; j < len; j++ {
	//		if nums[i] == nums[j] {
	//			return true
	//		}
	//	}
	//}
	//return false
	// todo：2 借用map与slice长度进行比较
	length := len(nums)
	tmpMap := make(map[int]int)
	for i := 0; i <= length - 1; i++ {
		tmpMap[nums[i]] = nums[i]
	}
	if length == len(tmpMap) {
		return false
	}
	return true
}



//存在重复元素
//给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
// 
//
//示例 1：
//
//输入：nums = [1,2,3,1]
//输出：true
//示例 2：
//
//输入：nums = [1,2,3,4]
//输出：false
//示例 3：
//
//输入：nums = [1,1,1,3,3,4,3,2,4,2]
//输出：true
//
//作者：力扣 (LeetCode)
//链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x248f5/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。