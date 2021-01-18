package main

// 131-分割回文串

//给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。 
//
// 返回 s 所有可能的分割方案。 
//
// 示例: 
//
// 输入: "aab"
//输出:
//[
//  ["aa","b"],
//  ["a","a","b"]
//] 
// Related Topics 深度优先搜索 动态规划 回溯算法 
// 👍 460 👎 0


func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	rs := [][]string{}
	var dfs func(string, int, []string)
	dfs = func(s string, index int, path []string) {
		if index == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			rs = append(rs, temp)
			return
		}

		for i := index; i < len(s); i++ {
			if !isPalindrome(s[index:i+1]) {
				continue
			}
			path = append(path, s[index:i+1])
			dfs(s, i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(s, 0, []string{})
	return rs
}

// 判断是否回文
func isPalindrome(b string) bool {
	for i:=0; i<len(b)-1-i; i++ {
		end := len(b)-1-i
		if b[i] != b[end] {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)

