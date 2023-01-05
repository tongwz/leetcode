package main

func singleNumber(nums []int) int {
	// todo: 1 效率低 执行慢 不太符合不适用额外空间
	//tmpMap := make(map[int]int)
	//for _, val := range nums {
	//	if v, ok := tmpMap[val]; ok {
	//		delete(tmpMap, v)
	//	} else {
	//		tmpMap[val] = val
	//	}
	//}
	//for val := range tmpMap {
	//	return val
	//}
	//return 0

	// todo: 2 还是用到了额外空间
	//tmpMap := make(map[int]int)
	//total := 0
	//halfTo := 0
	//for _, val := range nums {
	//	tmpMap[val] = val
	//	total += val
	//}
	//
	//for val := range tmpMap  {
	//	halfTo += val
	//}
	//halfTo = halfTo * 2
	//return halfTo - total

	//for index, val := range nums {
	//
	//}
	return 0
}

//只出现一次的数字
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//示例 1:
//
//输入: [2,2,1]
//输出: 1
//示例 2:
//
//输入: [4,1,2,1,2]
//输出: 4
//
//作者：力扣 (LeetCode)
//链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x21ib6/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
