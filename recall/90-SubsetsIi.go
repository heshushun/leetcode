package main

import "sort"

// 90-子集 II

//给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。 
//
// 说明：解集不能包含重复的子集。 
//
// 示例: 
//
// 输入: [1,2,2]
//输出:
//[
//  [2],
//  [1],
//  [1,2,2],
//  [2,2],
//  [1,2],
//  []
//] 
// Related Topics 数组 回溯算法 
// 👍 375 👎 0


func main() {
	subsetsWithDup([]int{1,2,2})
}

//leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	// data为子集
	var dfs func(nums []int, data []int, index int)
	sort.Ints(nums)
	dfs = func(nums []int, data []int, index int) {
		tmp := make([]int, len(data))
		copy(tmp,data)
		ans = append(ans, tmp)

		for i:=index; i<len(nums); i++ {
			if i>index && nums[i] == nums[i-1]{ // 剪枝
				continue
			}
			data = append(data, nums[i])
			dfs(nums, data, i+1)
			data = data[:len(data)-1]  // 回溯
		}
	}

	dfs(nums, []int{}, 0)
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)

