package main

import (
	"fmt"
)

/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/

func main() {
	s := "234"
	r := letterCombinations(s)
	fmt.Println(r)
}

var table = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}
	result := []string{""}

	// 广度优先
	for i:=0; i<len(digits); i++ {
		// 按键 key
		key := table[digits[i] - '0']
		var temp []string
		for j:=0; j<len(key); j++ {
			for z:=0; z<len(result); z++ {
				temp = append(temp, result[z] + string(key[j]))
			}
		}
		result = temp
	}
	return result
}