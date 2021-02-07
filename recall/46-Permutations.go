package main

// 46-全排列

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。 
//
// 示例: 
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//] 
// Related Topics 回溯算法 
// 👍 1115 👎 0


func main() {
	
}

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	res := [][]int{}
	var dfs func(nums []int, index int)
	dfs = func(nums []int, index int) {
		if index == len(nums){
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
		}

		for i:=index; i<len(nums); i++ {
			nums[i], nums[index] = nums[index], nums[i]
			dfs(nums, index+1)
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	dfs(nums, 0)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

