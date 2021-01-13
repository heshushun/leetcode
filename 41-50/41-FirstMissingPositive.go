package main

import (
	"fmt"
)

// 41-缺失的第一个正数

//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。 
//
// 
//
// 进阶：你可以实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案吗？ 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,0]
//输出：3
// 
//
// 示例 2： 
//
// 
//输入：nums = [3,4,-1,1]
//输出：2
// 
//
// 示例 3： 
//
// 
//输入：nums = [7,8,9,11,12]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 300 
// -231 <= nums[i] <= 231 - 1 
// 
// Related Topics 数组 
// 👍 914 👎 0


func main() {
	
}

//leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 将数组中所有小于等于 0 的数修改为 N+1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	// 将 没有被打标记的，并且 <=len(nums)的元素变为负数
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			fmt.Println(num -1)
			nums[num - 1] = -abs(nums[num -1])
		}
	}
	// 遍历找第一个正数
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
//leetcode submit region end(Prohibit modification and deletion)

