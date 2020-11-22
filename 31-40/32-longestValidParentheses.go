package main

/*
32. 最长有效括号
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/
func main() {
	l1 := "(()"
	r := longestValidParentheses(l1)
	println(r)
}


/*
动态规划
1.dp方程表示以i为截至的有效长度，"("对应为0
2.分两种情况分析："()","))"
3.记得追加前面的有效长度
*/
func longestValidParentheses(s string) int {
	var maxLength int
	dp := make([]int, len(s))

	for i:=1; i<len(s); i++ {
		if s[i-1] == '(' && s[i] == ')' {
			if i-2 > 0 {
				dp[i] = dp[i-2] + 2
			}else {
				dp[i] = 2
			}
		}

		if s[i-1] == ')' && s[i] == ')' {
			iLeft := i-1-dp[i-1]  // i对应的左括号位置
			if iLeft >= 0 && s[iLeft] == '(' {
				dp[i] = dp[i-1] + 2
				if iLeft -1 >= 0 {  // i对应的做好好位置的前一个
					dp[i] += dp[iLeft-1]  // 追加前面的有效长度
				}
			}
		}

		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}

	return maxLength
}