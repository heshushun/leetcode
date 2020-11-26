package main

// 34-在排序数组中查找元素的第一个和最后一个位置

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 你的算法时间复杂度必须是 O(log n) 级别。
//
// 如果数组中不存在目标值，返回 [-1, -1]。
//
// 示例 1:
//
// 输入: nums = [5,7,7,8,8,10], target = 8
//输出: [3,4]
//
// 示例 2:
//
// 输入: nums = [5,7,7,8,8,10], target = 6
//输出: [-1,-1]
// Related Topics 数组 二分查找
// 👍 670 👎 0


func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	left := -1
	right := -1

	if len(nums) == 0 || nums[len(nums) - 1] < target {
		return []int{left,right}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			left = i
		}

		j := len(nums)-i-1
		if nums[j] == target {
			right = j
		}
	}

	return []int{right, left}
}
//leetcode submit region end(Prohibit modification and deletion)

