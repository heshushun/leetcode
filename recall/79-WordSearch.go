package main

// 79-单词搜索

//给定一个二维网格和一个单词，找出该单词是否存在于网格中。 
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。 
//
// 
//
// 示例: 
//
// board =
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true
//给定 word = "SEE", 返回 true
//给定 word = "ABCB", 返回 false 
//
// 
//
// 提示： 
//
// 
// board 和 word 中只包含大写和小写英文字母。 
// 1 <= board.length <= 200 
// 1 <= board[i].length <= 200 
// 1 <= word.length <= 10^3 
// 
// Related Topics 数组 回溯算法 
// 👍 745 👎 0


func main() {
	
}

//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	n := len(board)    // 矩阵长度
	m := len(board[0]) // 矩阵宽度

	var dfs func(board [][]byte, word string, row, col, index int) bool

	dfs = func(board [][]byte, word string, row, col, index int) bool {
		// 找到了，不用继续往下找了
		if index == len(word) {
			return true
		}
		// 越界了
		if row < 0 || col < 0 || row == len(board) || col == len(board[row]) {
			return false
		}
		// 错误了
		if board[row][col] != word[index] {
			return false
		}

		// 重置它是为了回溯往回找的时候避免重复使用
		tmp := board[row][col]
		board[row][col] = byte(' ')
		// 向左
		if dfs(board, word, row-1, col, index+1) {
			return true
		}
		// 向右
		if dfs(board, word, row+1, col, index+1) {
			return true
		}
		// 向上
		if dfs(board, word, row, col-1, index+1) {
			return true
		}
		// 向下
		if dfs(board, word, row, col+1, index+1) {
			return true
		}
		// 结束了重置回来
		board[row][col] = tmp
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)

