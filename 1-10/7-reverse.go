package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
7. 整数反转
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
示例 1:
输入: 123
输出: 321

示例 2:
输入: -123
输出: -321

示例 3:
输入: 120
输出: 21
注意:
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

func main() {
	s := -839273298
	r := reverse(s)
	fmt.Println(r)
}

func reverse(x int) int {
	str := strconv.Itoa(x)
	sp := strings.Split(str, "")
	s := ""

	// 处理第一个 - 号
	if sp [0] == "-" {
		s += sp[0]
		sp = sp[1:]
	}
	// 反转
	for i:=len(sp)-1; i>=0; i-- {
		s += sp[i]
	}
	// string 转 int 自动去掉开头的0
	si,_ := strconv.Atoi(s)

	// 题目的范围
	if si >= math.MinInt32 && si <= math.MaxInt32 {
		return si
	}else{
		return 0
	}

}