package main

import "sync"

/*
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。
*/
func main() {
	l1 := "yyhellorrorldrrr"
	l2 := "hellorrorldrrryy"
	r := isAnagram(l1, l2)
	println(r)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 创建两个数组
	var wg sync.WaitGroup
	var wordstable_s = [26]int{}
	var wordstable_t = [26]int{}

	wg.Add(2)
	go func() {
		makewordtable(&wordstable_s, s)
		wg.Done()
	}()
	go func(){
		makewordtable(&wordstable_t, t)
		wg.Done()
	}()
	wg.Wait()

	// 最后比较一下两个数组
	return wordstable_s == wordstable_t
}

func makewordtable (wordtable *[26]int, s string){
	for i := 0; i<len(s); i++ {
		index := s[i] - 'a'
		wordtable[index] ++
	}
}