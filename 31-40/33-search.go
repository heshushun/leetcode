package main
/*
33. 搜索旋转排序数组
给你一个整数数组 nums ，和一个整数 target 。
该整数数组原本是按升序排列，但输入时在预先未知的某个点上进行了旋转。（例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] ）。
请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1
*/
func main() {
	nums := []int{4,5,6,7,0,1,2}
	target := 7
	r := search(nums, target)
	println(r)
}


// 二分查找
func search(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0

	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] { //左边有序
			if target > nums[mid] {
				l = mid + 1
			} else if target>=nums[l] {
				r = mid - 1
			}else{
				l=mid+1
			}
		} else { // 右侧有序
			if target < nums[mid] {
				r = mid - 1
			} else if target<=nums[r]{
				l = mid + 1
			}else{
				r=mid-1
			}
		}
	}
	return -1
}