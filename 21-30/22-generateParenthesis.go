package main

import (
	"fmt"
)

/*
   22. 括号生成
   数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
   示例：
   输入：n = 3
   输出：[
          "((()))",
          "(()())",
          "(())()",
          "()(())",
          "()()()"
        ]
*/
func main() {
s := 4
r := generateParenthesis(s)
fmt.Println(r)
}

var num int
var lBraceCnt, rBraceCnt int
func generateParenthesis(n int) []string {
	var res []string
	if n<=0 {
		return res
	}
	num = 2*n
	lBraceCnt, rBraceCnt = 0, 0
	backTrace(&res, []byte{}, 0)
	return res
}
// DFS（深度优先）+ 回溯算法
func backTrace(res *[]string, path []byte, idx int)  {
	if idx == num {
		*res = append(*res, string(path))
		return
	}

	if lBraceCnt < num/2 {
		if len(path) > idx {
			path[idx] = '('
		}else {
			path = append(path, '(')
		}
		lBraceCnt++
		backTrace(res, path, idx+1)
		lBraceCnt--
	}
	if rBraceCnt < lBraceCnt {
		if len(path) > idx {
			path[idx] = ')'
		}else {
			path = append(path, ')')
		}
		rBraceCnt++
		backTrace(res, path, idx+1)
		rBraceCnt--
	}
}