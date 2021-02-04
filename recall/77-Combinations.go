package main

// 77-组合

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。 
//
// 示例: 
//
// 输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//] 
// Related Topics 回溯算法 
// 👍 481 👎 0


func main() {
	
}

//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	res := [][]int{}
	var dfs func([]int, int)
	dfs = func(nums []int, index int) {
		if len(nums) == k {
			tmp := make([]int, k)
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}

		for i := index; i<=n; i++ {
			if n-i+1 < k-len(nums){  // 剪枝
				break
			}
			nums = append(nums, i)
			dfs(nums, i+1)
			nums = nums[:len(nums)-1]  // 回溯
		}
	}
	dfs([]int{}, 1)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

