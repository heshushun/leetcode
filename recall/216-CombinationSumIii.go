package main

// 216-组合总和 III

//找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。 
//
// 说明： 
//
// 
// 所有数字都是正整数。 
// 解集不能包含重复的组合。 
// 
//
// 示例 1: 
//
// 输入: k = 3, n = 7
//输出: [[1,2,4]]
// 
//
// 示例 2: 
//
// 输入: k = 3, n = 9
//输出: [[1,2,6], [1,3,5], [2,3,4]]
// 
// Related Topics 数组 回溯算法 
// 👍 252 👎 0


func main() {
	
}

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum3(k int, n int) [][]int {
	/*
	   回溯 + 剪枝
	   剪枝条件： > n
	   去重条件： 加 index 控制
	*/
	ans := make([][]int, 0) // 结果
	temp := make([]int, 0)

	var dfs func(k int, n int, start int)

	dfs = func(k int, n int, index int) {
		if n == 0 && k == 0 {
			tmp := make([]int, len(temp))
			copy(tmp, temp)
			ans = append(ans, tmp)
			return
		}

		for i:=index; i<=9; i++ {
			if n-i < 0 {  // 剪枝（不能超过n）
				break
			}
			if k == 1 && n-i > 0 {  // 剪枝 （只剩下一个数 但是还不够）
				continue
			}
			temp = append(temp, i)
			dfs(k-1, n-i, i+1)
			temp = temp[:len(temp)-1]  // 回溯
		}
	}
	dfs(k, n, 1)
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)

