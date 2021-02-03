package main

// 47-全排列 II

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 8 
// -10 <= nums[i] <= 10 
// 
// Related Topics 回溯算法 
// 👍 574 👎 0


func main() {
	
}

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	var dfs func(nums []int, index int)
	dfs = func(nums []int, index int) {
		if index == len(nums){
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
		}
		// 每次递归中维护一个 map 来检查是否在此阶段将两个相同的数交换了位置，若交换了则必定有所重复。
		visit := map[int]int{}
		for i:=index; i<len(nums); i++ {
			if _, ok := visit[nums[i]]; ok {
				continue
			}
			nums[i], nums[index] = nums[index], nums[i]

			dfs(nums, index+1)
			nums[i], nums[index] = nums[index], nums[i]
			visit[nums[i]] = 1
		}
	}
	dfs(nums, 0)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

