package main
/*
27. 移除元素
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1:
给定 nums = [3,2,2,3], val = 3,
函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
你不需要考虑数组中超出新长度后面的元素。
示例 2:
给定 nums = [0,1,2,2,3,0,4,2], val = 2,
函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
注意这五个元素可为任意顺序。
你不需要考虑数组中超出新长度后面的元素。
*/
func main() {
	l := []int{1, 2, 2, 3, 4, 5, 5}
	r := removeElement(l, 3)
	println(r)
}

// func removeElement(nums []int, val int) int {
//     newLen := 0
//     for i:= 0; i < len(nums); i++ {
//         num := nums[i]
//         if num != val {
//             nums[newLen] = num
//             newLen ++
//         }
//     }
//     return newLen
// }

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		num := nums[i]
		if num == val {
			nums = append(nums[:i], nums[i+1:]...)
		}else {
			i++
		}
	}
	return len(nums)
}