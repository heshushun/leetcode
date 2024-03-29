package main

import (
	"fmt"
)

/*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
输入: "()"
输出: true

示例 2:
输入: "()[]{}"
输出: true

示例 3:
输入: "(]"
输出: false

示例 4:
输入: "([)]"
输出: false
*/

func main() {
	s := "()[]{}"
	r := isValid(s)
	fmt.Println(r)
}

// 栈的方式
func isValid(s string) bool {
	var stack = make([]byte, 0)
	var bracketsMap = map[byte]byte {
		')': '(',
		']': '[',
		'}': '{',
	}

	for i:=0; i<len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		value, ok := bracketsMap[s[i]]
		if ok {
			if value == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}else {
				return false
			}
		}
	}
	return len(stack) == 0
}