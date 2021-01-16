package main

import (
	"strconv"
)

// 401-二进制手表

//二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。 
//
// 每个 LED 代表一个 0 或 1，最低位在右侧。 
//
// 
//
// 例如，上面的二进制手表读取 “3:25”。 
//
// 给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。 
//
// 
//
// 示例： 
//
// 输入: n = 1
//返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "
//0:32"] 
//
// 
//
// 提示： 
//
// 
// 输出的顺序没有要求。 
// 小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。 
// 分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。 
// 超过表示范围（小时 0-11，分钟 0-59）的数据将会被舍弃，也就是说不会出现 "13:00", "0:61" 等时间。 
// 
// Related Topics 位运算 回溯算法 
// 👍 213 👎 0


func main() {
	
}

//leetcode submit region begin(Prohibit modification and deletion)
//1.通过回溯法，当亮灯数==num，判断小时和分是否符合规则，
// 符合则添加到[]string里，不符合就return
//2.回溯的时候，当下标小于4的时候，数值添加到小时，大于4小于10的时候，
// 数值添加到分，大于等于10的时候，return
func readBinaryWatch(num int) []string {
	solutions := []string{}
	lists := []int{1, 2, 4, 8, 1, 2, 4, 8, 16, 32}
	var time string
	var backstrack func(int, int, int, int)

	backstrack = func(index int, cnt int, hour int, minute int) {
		if cnt == num {
			if minute<0 || minute>59 || hour<0 || hour>11 {
				return
			}
			if minute<10 {
				time = strconv.Itoa(hour)+":0"+strconv.Itoa(minute)
			}else{
				time = strconv.Itoa(hour)+":"+strconv.Itoa(minute)
			}
			solutions = append(solutions, time)
			return
		}
		for index<10 {
			if index<=3 {
				hour+=lists[index]
				backstrack(index+1, cnt+1, hour, minute)
				hour-=lists[index]
			}else {
				minute+=lists[index]
				backstrack(index+1, cnt+1, hour, minute)
				minute-=lists[index]
			}
			index++
		}
		return
	}

	backstrack(0,0,0,0)
	return solutions
}
//leetcode submit region end(Prohibit modification and deletion)

