package main

import "fmt"

/*
5.最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
示例 1：
输入: "babad"
输出: "bab"
示例 2：
输入: "cbbd"
输出: "bb"
*/

func main() {
	s := "babaaaaadaaaaaa"
	r := longestPalindrome(s)
	fmt.Println(r)
}

func longestPalindrome(s string) string {

	if len(s) == 0{
		return s
	}
	length := len(s)
	longestLeft := 0
	longestRight := 0

	for i := 0; i < length; i++ {
		// bab类型
		left, right := expandAroundCenter(i, i, s)
		if (right - left) > (longestRight - longestLeft){
			longestLeft = left
			longestRight = right
		}
		// bb类型
		left, right = expandAroundCenter(i, i+1, s)
		if (right - left) > (longestRight - longestLeft){
			longestLeft = left
			longestRight = right
		}
	}
	return s[longestLeft:longestRight+1]
}

func expandAroundCenter (left int, right int, s string) (leftEnd, rightEnd int){

	if left != right{
		leftEnd, rightEnd = 0, 0
	}else{
		leftEnd, rightEnd = left, right
	}

	for {
		if (left < 0 || right >= len(s)){
			return leftEnd, rightEnd
		}
		if s[left] != s[right]{
			return leftEnd, rightEnd
		}
		leftEnd = left
		rightEnd = right
		left --
		right ++
	}
	return leftEnd, rightEnd
}