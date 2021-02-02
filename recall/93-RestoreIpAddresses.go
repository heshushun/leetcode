package main

import (
	"strconv"
	"strings"
)

// 93-复原IP地址

//给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。 
//
// 有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。 
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
// 和 "192.168@1.1" 是 无效的 IP 地址。 
//
// 
//
// 示例 1： 
//
// 输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
// 
//
// 示例 2： 
//
// 输入：s = "0000"
//输出：["0.0.0.0"]
// 
//
// 示例 3： 
//
// 输入：s = "1111"
//输出：["1.1.1.1"]
// 
//
// 示例 4： 
//
// 输入：s = "010010"
//输出：["0.10.0.10","0.100.1.0"]
// 
//
// 示例 5： 
//
// 输入：s = "101023"
//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length <= 3000 
// s 仅由数字组成 
// 
// Related Topics 字符串 回溯算法 
// 👍 492 👎 0


func main() {
	
}

//leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	var dfs func(data []string, index int)

	dfs = func(data []string, index int) {
		// 终止条件 最大4段
		if len(data) == 4 && index == len(s) {
			ans = append(ans, strings.Join(data, "."))
			return
		}
		// 不符合条件
		if len(data) == 4 && index < len(s) {
			return
		}

		// 每一段最大选择3位
		for i := 1; i <= 3; i++ {
			if index+i-1 >= len(s){  // 剪枝 溢出了
				return
			}
			if i != 1 && s[index] == '0' {  // 剪枝 第一位开头不能是0
				return
			}
			str := s[index: index+i]
			if n, _ := strconv.Atoi(str); n > 255 { // 剪枝 值不能大于255
				return
			}
			data = append(data, str)
			dfs(data, index+i)
			data = data[:len(data)-1]  // 回溯
		}
	}
	dfs([]string{}, 0)
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)

