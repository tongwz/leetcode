package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	var tmpMap = make(map[int]int)
	var res []int

	for _, val := range nums1 {
		if _, ok := tmpMap[val]; ok {
			tmpMap[val]++
		} else {
			tmpMap[val] = 1
		}
	}
	for _, val := range nums2 {
		if _, ok := tmpMap[val]; ok {
			tmpMap[val]--
			if tmpMap[val] >= 0 {
				res = append(res, val)
			}
		}
	}
	return res
}

func main() {
	var nums1 = []int{1, 2, 2, 1}
	var nums2 = []int{2, 2}
	re := intersect(nums1, nums2)
	fmt.Printf("获取到的结果是 %#v \n", re)
}

//给你两个整数数组nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
//
//
//
//示例 1：
//
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
//示例 2:
//
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[4,9]
//
//
//提示：
//
//1 <= nums1.length, nums2.length <= 1000
//0 <= nums1[i], nums2[i] <= 1000
//
//
//进阶：
//
//如果给定的数组已经排好序呢？你将如何优化你的算法？
//如果 nums1 的大小比 nums2 小，哪种方法更优？
//如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
//
//作者：力扣 (LeetCode)
//链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2y0c2/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
