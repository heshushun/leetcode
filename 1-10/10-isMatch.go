package main

import (
	"fmt"
)

/*
10. 正则表达式匹配
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
说明:
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

示例 1:
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。

示例 2:
输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

示例 3:
输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
*/

func main() {
	s := "839273298"
	r := isMatch(s, "8.*")
	fmt.Println(r)
}

func isMatch(s string, p string) bool {
	dp := make([][]int, len(s)+1)
	for i:=0; i<=len(s); i++ {
		dp[i] = make([]int, len(p)+1)
	}
	return isMatchCore(s, p, 0, 0, dp)

}

func isMatchCore (s, p string, i, j int, dp [][]int) bool{
	if dp[i][j] == 1 {
		return true
	}else if dp[i][j] == -1{
		return false
	}

	flag := false
	if j == len(p) {
		if i == len(s) {
			flag = true
		} else {
			flag = false
		}
	}

	if j == len(p)-1 {
		if i == len(s) {
			flag = false
		} else if s[i] == p[j] || p[j] == '.' {
			flag = isMatchCore(s, p, i+1, j+1, dp)
		}else{
			flag = false
		}
	}

	if j < len(p) -1 {
		if p[j+1] == '*' {
			if i<= len(s) -1 && (s[i] == p[j] || p[j] == '.') {
				flag = isMatchCore(s, p, i+1, j+2, dp) || isMatchCore(s, p, i+1, j, dp) || isMatchCore(s, p, i, j+2, dp)
			}else {
				flag = isMatchCore(s, p, i, j+2, dp)
			}
		} else if i <= len(s) -1 && (s[i] == p[j] || p[j] == '.') {
			flag = isMatchCore(s, p, i+1, j+1, dp)
		}
	}

	if flag == true {
		dp[i][j] = 1
	}else {
		dp[i][j] = -1
	}
	return flag
}