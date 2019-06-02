package main

/*
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

思路: 如果 s[j]在 [i, j) 范围内有与 j 重复的字符，我们不需要逐渐增加 i 。
我们可以直接跳过 [i，j] 范围内的所有元素，并将 i 变为 j + 1。
*/
import "fmt"

func main() {
	s := "alqebriavxoo"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	subMap := make(map[byte]int)
	for i, j := 0, 0; j < n; j++ {
		if v, ok := subMap[byte(s[j])]; ok {
			if i < v {
				i = v
			}
		}
		if ans < (j - i + 1) {
			ans = j - i + 1
		}
		subMap[byte(s[j])] = j + 1
	}
	return ans
}
