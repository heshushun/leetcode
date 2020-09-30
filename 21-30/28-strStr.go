package main
/*
28. 实现 strStr()
实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
示例 1:
输入: haystack = "hello", needle = "ll"
输出: 2

示例 2:
输入: haystack = "aaaaa", needle = "bba"
输出: -1
*/
func main() {
	l1 := "helloWorld"
	l2 := "ll"
	r := strStr(l1, l2)
	println(r)
}

func strStr(haystack string, needle string) int {
	len1 := len(haystack)
	len2 := len(needle)

	if len2 == 0 {
		return 0
	}
	if len1 == 0 || len1 < len2 {
		return -1
	}

	for i := 0; i <= len1 - len2; i++ {
		if haystack[i : i+len2] == needle {
			return i
		}
	}
	return -1
}
