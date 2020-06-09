package main

import "fmt"

/*
6. Z 字形变换
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
请你实现这个将字符串进行指定行数变换的函数：
string convert(string s, int numRows);
示例:
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:
L     D     R
E   O E   I I
E C   I H   N
T     S     G
*/

func main() {
	s := "0123456789ABCDEFGH"
	r := convert(s, 5)
	fmt.Println(r)
}

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	var rst string
	// 等差数列差值
	step := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += step {
			rst += string(s[j])
			tmp := j + step - 2*i
			//处理不是首行和最后一行,加入两个满列之间的值
			if i != 0 && i != numRows - 1 && tmp < len(s){
				rst += string(s[tmp])
			}
		}
	}
	return rst
}