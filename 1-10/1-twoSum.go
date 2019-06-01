package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

import "fmt"

func main() {
	nums := []int {1,3,4,6,7,10}
	target := 17
	newArr := twoSum(nums,target)
	fmt.Println(newArr)
}

func twoSum(nums []int, target int) []int{
	maps := make(map[int] int)
	for index, value := range nums{
		complement := target - value
		if _, ok := maps[complement];ok {
			return []int{maps[complement],index}
		} else{
			maps[value] = index
		}
	}
	return []int{}
}