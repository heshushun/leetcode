package main

import (
	"sort"
)

// 1079-活字印刷

//你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。 
//
// 注意：本题中，每个活字字模只能使用一次。 
//
// 
//
// 示例 1： 
//
// 输入："AAB"
//输出：8
//解释：可能的序列为 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"。
// 
//
// 示例 2： 
//
// 输入："AAABBC"
//输出：188
// 
//
// 
//
// 提示： 
//
// 
// 1 <= tiles.length <= 7 
// tiles 由大写英文字母组成 
// 
// Related Topics 回溯算法 
// 👍 101 👎 0


func main() {
	numTilePossibilities("AAB")
}

//leetcode submit region begin(Prohibit modification and deletion)
func numTilePossibilities(tiles string) int {
	var count int
	var dfs func([]byte, []bool)

	// 排序 方便去重
	byTiles := []byte(tiles)
	sort.Slice(byTiles, func(i, j int) bool {
		return byTiles[i] < byTiles[j]
	})

	// visit 每轮已选择数
	dfs = func(data []byte, visit []bool) {
		// 没有可选择的了，终止
		if len(data) == len(tiles){
			return
		}

		// 每次都从头开始
		for i:=0; i < len(tiles); i++ {
			if visit[i] {  // 剪枝
				continue
			}
			if i>0 && byTiles[i]==byTiles[i-1] && !visit[i-1] {  // 剪枝  !visit[i-1]上一个如果这轮没有选择，选后面那个肯定会跟上轮有结果重复
				continue
			}
			count++
			visit[i] = true
			dfs(append(data, byTiles[i]), visit)
			visit[i] = false  // 回溯
		}
	}
	dfs([]byte{}, make([]bool, len(tiles)))
	return count
}
//leetcode submit region end(Prohibit modification and deletion)

