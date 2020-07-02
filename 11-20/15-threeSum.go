package main

import "fmt"

/*
15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组
注意：答案中不可以包含重复的三元组。

示例：
给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func main() {
	s := []int{-1, 0, 1, 2, -1, -4}
	r := threeSum(s)
	fmt.Println(r)
}

/***
特判，对于数组长度 nn，如果数组为 nullnull 或者数组长度小于 33，返回 [][]。
对数组进行排序。
遍历排序后数组：
1.若 nums[i]>0nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 00，直接返回结果。
2.对于重复元素：跳过，避免出现重复解
3.令左指针 L=i+1L=i+1，右指针 R=n-1R=n−1，当 L<RL<R 时，执行循环：
    当 nums[i]+nums[L]+nums[R]==0nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,RL,R 移到下一位置，寻找新的解
    若和大于 00，说明 nums[R]nums[R] 太大，RR 左移
    若和小于 00，说明 nums[L]nums[L] 太小，LL 右移
***/

func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return [][]int{}
	}

	QuickSort(nums)

	ret := make([][]int, 0, 0)
	for i := 0; i < len(nums) -1; i++ {
		l, r := i + 1, len(nums) - 1

		// 不可能情况
		if nums[i] > 0 || nums[i] + nums[l] > 0 {
			break
		}
		// i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 双指针方式
		for l < r {
			// l 去重
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			// r 去重
			if r < len(nums)-2 && nums[r] == nums[r+1] {
				r--
				continue
			}

			if nums[i] + nums[l] + nums[r] > 0{
				r--
			}else if nums[i] + nums[l] + nums[r] < 0{
				l++
			}else{
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
		}
	}
	return ret
}

func QuickSort(array []int) {
	if len(array) < 2 { return }

	pivot := array[0]
	l, r := 0, len(array) - 1

	for i := 1; i <= r; {
		if array[i] < pivot {
			array[l], array[i] = array[i], array[l]
			l++; i++
		} else {
			array[r], array[i] = array[i], array[r]
			r--
		}
	}

	QuickSort(array[:l])
	QuickSort(array[l + 1:])
}
