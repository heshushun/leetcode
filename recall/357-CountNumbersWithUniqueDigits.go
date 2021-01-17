package main

// 357-计算各个位数不同的数字个数

//给定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10n 。 
//
// 示例: 
//
// 输入: 2
//输出: 91 
//解释: 答案应为除去 11,22,33,44,55,66,77,88,99 外，在 [0,100) 区间内的所有数字。
// 
// Related Topics 数学 动态规划 回溯算法 
// 👍 110 👎 0


func main() {
	
}

//leetcode submit region begin(Prohibit modification and deletion)
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	used := make([]bool, 10)
	var dfs func(int, int, []bool) int

	dfs = func(n int, index int, used []bool) int {
		if index == n {
			return 0
		}
		result := 0
		for num := 0; num <= 9; num++ {
			// 剪枝： 多位数时，第一个数字不能为0
			if n >= 2 && index == 1 && num == 0 {
				continue
			}
			// 剪枝：不能使用用过的数字
			if used[num] {
				continue
			}
			used[num] = true
			result += dfs(n, index+1, used) + 1  // 加上当前的数字
			used[num] = false
		}
		return result
	}
	return dfs(n, 0, used)
}
//leetcode submit region end(Prohibit modification and deletion)

