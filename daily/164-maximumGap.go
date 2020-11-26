package main

// 164-最大间距

//给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
//
// 如果数组元素个数小于 2，则返回 0。
//
// 示例 1:
//
// 输入: [3,6,9,1]
//输出: 3
//解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
//
// 示例 2:
//
// 输入: [10]
//输出: 0
//解释: 数组元素个数小于 2，因此返回 0。
//
// 说明:
//
//
// 你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
// 请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。
//
// Related Topics 排序
// 👍 260 👎 0


func main() {
	para := []int{1, 3, 9, 7, 4}
	res := maximumGap(para)
	print(res)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maximumGap(nums []int) int {
	if len(nums) < 2{
		return 0
	}
	nums = quickSort(nums)
	a := 0
	for i := 1; i < len(nums); i++{
		if nums[i]-nums[i-1] > a{
			a = nums[i] - nums[i-1]
		}
	}
	return a
}

func quickSort(a []int) []int{
	return _quickSort(a, 0, len(a)-1)
}
func _quickSort(a []int, l, r int) []int{
	if l < r{
		partitionIndex := partition(a, l, r)
		_quickSort(a, l, partitionIndex-1)
		_quickSort(a, partitionIndex+1, r)
	}
	return a
}
func partition(a []int, l, r int) int{
	pivot := l
	index := pivot+1
	for i := index; i <= r; i++{
		if a[pivot] > a[i]{
			swap(a, i, index)
			index++
		}
	}
	swap(a, pivot, index-1)
	return index-1
}
func swap(a []int, i, j int){
	a[i], a[j] = a[j], a[i]
}
//leetcode submit region end(Prohibit modification and deletion)

