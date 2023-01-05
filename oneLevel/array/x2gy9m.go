package main

import "fmt"

/**
 * @note: 删除排序数组中的重复项
 * @auth: tongWz
 * @date: 2022年6月20日16:15:02
**/
func removeDuplicates(nums []int) int {
	fmt.Printf("arr的指针2是：%p \n", &nums)
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func main() {
	var arr = []int{0, 0, 1, 1, 2}
	fmt.Printf("arr的指针是：%p \n", &arr)
	// 这里多一个返回值因为 切片内容修改后只会修改函数内的切片 并不会影响函数外的切片
	last := removeDuplicates(arr)
	fmt.Printf("获取到的新长度为：%d, arr现在值：%+v, 指针：%p \n", last, arr, &arr)
}

//给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
//
//由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。
//
//将最终结果插入 nums 的前 k 个位置后返回 k 。
//
//不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

//判题标准:
//
//系统会用下面的代码来测试你的题解:
//
//int[] nums = [...]; // 输入数组
//int[] expectedNums = [...]; // 长度正确的期望答案
//
//int k = removeDuplicates(nums); // 调用
//
//assert k == expectedNums.length;
//for (int i = 0; i < k; i++) {
//assert nums[i] == expectedNums[i];
//}
//如果所有断言都通过，那么您的题解将被 通过。
//作者：力扣 (LeetCode)
//链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2gy9m/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
