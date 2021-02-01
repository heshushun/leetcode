package main

import "sort"

// 78-子集

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。 
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [0]
//输出：[[],[0]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10 
// -10 <= nums[i] <= 10 
// nums 中的所有元素 互不相同 
// 
// Related Topics 位运算 数组 回溯算法 
// 👍 968 👎 0


func main() {
	
}

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	// data为子集
	var dfs func(nums []int, data []int, index int)
	sort.Ints(nums)
	dfs = func(nums []int, data []int, index int) {
		tmp := make([]int, len(data))
		copy(tmp,data)
		ans = append(ans, tmp)

		for i:=index; i<len(nums); i++ {
			data = append(data, nums[i])
			dfs(nums, data, i+1)
			data = data[:len(data)-1]  // 回溯
		}
	}

	dfs(nums, []int{}, 0)
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)

