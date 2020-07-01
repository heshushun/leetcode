package main

import "fmt"

/*
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例 1:
输入: ["flower","flow","flight"]
输出: "fl"

示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:
所有输入只包含小写字母 a-z 。
*/

func main() {
	s := []string{"flower","flow","flight"}
	r := longestCommonPrefix(s)
	fmt.Println(r)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minIndex := len(strs[0])-1
	// 以第一个为基准 比较后面每个string
	for i:=0; i < len(strs); i++ {
		k:=0  // 动态的索引
		if strs[0][k] != strs[i][k] {
			return ""
		}
		for ; k <= minIndex; k++ {
			if len(strs[i])-1 < k || strs[0][k] != strs[i][k]{
				break
			}
		}
		minIndex = k-1
	}
	s := strs[0][:minIndex+1]
	return s
}
