package main

// 767-重构字符串

//给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。
//
// 若可行，输出任意可行的结果。若不可行，返回空字符串。
//
// 示例 1:
//
//
//输入: S = "aab"
//输出: "aba"
//
//
// 示例 2:
//
//
//输入: S = "aaab"
//输出: ""
//
//
// 注意:
//
//
// S 只包含小写字母并且长度在[1, 500]区间内。
//
// Related Topics 堆 贪心算法 排序 字符串
// 👍 223 👎 0


func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)

//双指针找不同的字符，如果字符不同，快慢指针一起后移；
//如果字符相同，慢指针指向快指针位置，快指针往后移直到找到不同的字符，然后交换位置，慢指针后移一位，快指针回到慢指针+1的位置
//字符串反转再来一遍，返回结果。
func reorganizeString(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	var arr = []byte(s)
	ans := reorg(arr, n)
	if len(ans) > 0 {
		return ans
	}
	// 字符串反转
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	// 再来一遍
	res := reorg(arr, n)
	return res
}

func reorg(arr []byte, n int) string {
	var slow, fast = 0, 1
	for fast < n {
		if arr[slow] != arr[fast] {
			slow++
			fast++
			continue
		}
		slow = fast
		for fast < n && arr[slow] == arr[fast] {
			fast++
		}
		if fast == n {
			return ""
		}
		arr[slow], arr[fast] = arr[fast], arr[slow]
		slow++
		fast = slow + 1
	}
	return string(arr)
}
//leetcode submit region end(Prohibit modification and deletion)

