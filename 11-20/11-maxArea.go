package _1_20

import "fmt"

/*
11. 盛最多水的容器
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
*/
func main() {
	s := []int{1,8,6,2,5,4,8,3,7}
	r := maxArea(s)
	fmt.Println(r)
}

func maxArea(height []int) int {

	left, right, max := 0, len(height)-1, 0

	for left < right {
		capValue := min(height[left], height[right]) * (right - left)
		if capValue > max {
			max = capValue
		}
		if height[left] < height[right]{
			left ++
		}else {
			right --
		}
	}
	return max
}

func min(a int, b int) int{
	if a > b {
		return b
	}
	return a
}