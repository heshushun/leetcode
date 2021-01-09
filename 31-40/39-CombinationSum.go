package main

import (
	"sort"
)

// 39-组合总和

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。 
//
// candidates 中的数字可以无限制重复被选取。 
//
// 说明： 
//
// 
// 所有数字（包括 target）都是正整数。 
// 解集不能包含重复的组合。 
// 
//
// 示例 1： 
//
// 输入：candidates = [2,3,6,7], target = 7,
//所求解集为：
//[
//  [7],
//  [2,2,3]
//]
// 
//
// 示例 2： 
//
// 输入：candidates = [2,3,5], target = 8,
//所求解集为：
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//] 
//
// 
//
// 提示： 
//
// 
// 1 <= candidates.length <= 30 
// 1 <= candidates[i] <= 200 
// candidate 中的每个元素都是独一无二的。 
// 1 <= target <= 500 
// 
// Related Topics 数组 回溯算法 
// 👍 1110 👎 0


func main() {
	
}

//leetcode submit region begin(Prohibit modification and deletion)
// result 定义为全局变量，减少回溯函数的一个参数传入量，看起来简洁点
var result [][] int
func combinationSum(candidates []int, target int) [][]int {
	// 排序数组，才能使用剪枝功能
	sort.Ints(candidates)
	// clear result
	// 如果这里不清理result, 那么每次结果就会累计起来
	result = result[0:0]
	dfs(target, candidates, []int(nil))

	return result
}

//确定回溯函数返回值以及参数
//确定回溯函数终止条件
//确定回溯函数的遍历过程

// 1.确定参数和返回值
func dfs(sum int, candidates []int, path []int)  {
	// 2.确定终止条件
	if sum == 0 {
		result = append(result, append([]int(nil), path...))
	}

	// 3.遍历当前层中所有的元素
	for i, num := range candidates {
		if sum - num  < 0{
			return
		}
		// append(path, num) 将会创建一个新的path切片
		// 因此当前层的Path不会改变，所以自带回溯效果
		dfs(sum - num, candidates[i:], append(path, num))
	}
}
//leetcode submit region end(Prohibit modification and deletion)

