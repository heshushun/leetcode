package main

/*
763. 划分字母区间
字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一个字母只会出现在其中的一个片段。返回一个表示每个字符串片段的长度的列表。

示例 1：
输入：S = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca", "defegde", "hijhklij"。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
 
提示：
S的长度在[1, 500]之间。
S只包含小写字母 'a' 到 'z' 。
*/
func main() {
	l1 := "ababcbacadefegdehijhklij"
	r := partitionLabels(l1)
	println(r)
}

func partitionLabels(S string) []int {
	// lastIndex 记录每一个字符最后出现的位置
	lastIndex := make([]int, 26)
	for i, c := range  S {
		lastIndex[int(c-'a')] = i
	}
	// 遍历找这个位置 是不是这个字母最后的位置，是就说明这就是一个分段
	ans := make([]int, 0)
	start := 0
	end := 0
	for i := 0; i < len(S); i++ {
		end = max(lastIndex[int(S[i] - 'a')], end)
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}