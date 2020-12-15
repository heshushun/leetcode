package main

import "strconv"

// 738-单调递增的数字

//给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。 
//
// （当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。） 
//
// 示例 1: 
//
// 输入: N = 10
//输出: 9
// 
//
// 示例 2: 
//
// 输入: N = 1234
//输出: 1234
// 
//
// 示例 3: 
//
// 输入: N = 332
//输出: 299
// 
//
// 说明: N 是在 [0, 10^9] 范围内的一个整数。 
// Related Topics 贪心算法 
// 👍 106 👎 0


func main() {
	res := monotoneIncreasingDigits(896)
	println(res)
}

//leetcode submit region begin(Prohibit modification and deletion)
func monotoneIncreasingDigits(N int) int {
	// 先转化位字节数组
	s := strconv.Itoa(N)
	ss := []byte(s)
	n := len(ss)
	// 如果只有一位数直接返回
	if n == 1{
		return N
	}
	// 贪心算法 从右往左遍历
	for i:=n-2; i>=0; i-- {
		if ss[i] > ss[i+1]{
			ss[i] -= 1
			// 后置位要全部置为9
			for j:=i+1; j<n; j++{
				ss[j] = '9'
			}
		}
	}
	// 数组再转为字符串和数字
	str := string(ss)
	res, _ := strconv.Atoi(str)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

