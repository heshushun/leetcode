package main

// 321-拼接最大数

//给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，表示两个自然数各位上的数字。现在从这两个数组中选出 k (k <= m + n) 个数字拼接
//成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序。
//
// 求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。
//
// 说明: 请尽可能地优化你算法的时间和空间复杂度。
//
// 示例 1:
//
// 输入:
//nums1 = [3, 4, 6, 5]
//nums2 = [9, 1, 2, 5, 8, 3]
//k = 5
//输出:
//[9, 8, 6, 5, 3]
//
// 示例 2:
//
// 输入:
//nums1 = [6, 7]
//nums2 = [6, 0, 4]
//k = 5
//输出:
//[6, 7, 6, 0, 4]
//
// 示例 3:
//
// 输入:
//nums1 = [3, 9]
//nums2 = [8, 9]
//k = 3
//输出:
//[9, 8, 9]
// Related Topics 贪心算法 动态规划
// 👍 249 👎 0


func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubsequence(a []int, k int) (s []int) {
	for i, v := range a {
		for len(s) > 0 && len(s)+len(a)-1-i >= k && v > s[len(s)-1] {
			s = s[:len(s)-1]
		}
		if len(s) < k {
			s = append(s, v)
		}
	}
	return
}

func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	for i := range merged {
		if lexicographicalLess(a, b) {
			merged[i], b = b[0], b[1:]
		} else {
			merged[i], a = a[0], a[1:]
		}
	}
	return merged
}

func maxNumber(nums1, nums2 []int, k int) (res []int) {
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}
	for i := start; i <= k && i <= len(nums1); i++ {
		s1 := maxSubsequence(nums1, i)
		s2 := maxSubsequence(nums2, k-i)
		merged := merge(s1, s2)
		if lexicographicalLess(res, merged) {
			res = merged
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

