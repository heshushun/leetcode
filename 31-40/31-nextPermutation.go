package main

/*
31. 下一个排列
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
func main() {
	l1 := []int{1,2,3,4,5,6}
	nextPermutation(l1)
}

func nextPermutation(nums []int)  {
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2 , len(nums)-1, len(nums)-1
	//find A[i]<A[j]
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		// find A[i] < A[k]
		for nums[i] >= nums[k] {
			k--
		}

		// swap A[i], A[k]
		nums[i], nums[k] = nums[k], nums[i]
	}

	// reverse A[j:end]
	for i, j := j, len(nums)-1; i<j; i,j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
