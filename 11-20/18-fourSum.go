package main

import (
	"fmt"
	"sort"
)

/*
18. 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
注意：
答案中不可以包含重复的四元组。

示例：
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func main() {
	s := []int{-1, 0, 1, 2, -1, -4}
	r := fourSum(s, 1)
	fmt.Println(r)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	nSum(nums, target, 0, 4, &res, []int{})
	return res
}

func nSum(nums []int, target int, pos int, n int, res *[][]int, cur []int) {
	if n == 2 {
		left, right := pos, len(nums) - 1
		for left < right {
			if sum := nums[left] + nums[right]; sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				t := make([]int, len(cur) + 2)
				copy(t, cur)
				t[len(t) - 2] = nums[left]
				t[len(t) - 1] = nums[right]
				*res = append(*res, t)
				left++
				right--
				for left < right && nums[left] == nums[left - 1] {
					left++
				}
				for left < right && nums[right] == nums[right + 1] {
					right--
				}
			}
		}
		return
	}

	for i := pos; i < len(nums) - n + 1; i++ {
		if target < nums[i] * n || target > nums[len(nums)-1] * n {
			break
		}
		if i > pos && nums[i] == nums[i-1] {
			continue
		}
		cur = append(cur, nums[i])
		nSum(nums, target - nums[i], i + 1, n - 1, res, cur)
		cur = cur[:len(cur)-1]
	}
}



