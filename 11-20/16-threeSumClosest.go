package main

import (
	"fmt"
	"sort"
)

/*
16. 最接近的三数之和
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
示例：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
*/

func main() {
	s := []int{-1, 0, 1, 2, -1, -4}
	r := threeSumClosest(s, 1)
	fmt.Println(r)
}

func threeSumClosest(nums []int, target int) int {

	// 排序
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		l, r := i + 1, len(nums) - 1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			// 终止条件
			if sum == target {
				return target
			}
			// 更新 result
			if abs(sum-target) < abs(result-target){
				result = sum
			}
			// 移动方向
			if sum < target {
				l ++
			} else{
				r --
			}
		}
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

