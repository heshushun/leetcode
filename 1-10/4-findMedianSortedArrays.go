package main

import "fmt"

/*
4.寻找两个有序数组的中位数
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。

示例:
nums1 = [1, 3]  nums2 = [2]
则中位数是 2.0

nums1 = [1, 2]  nums2 = [3, 4]
则中位数是 (2 + 3)/2 = 2.5
*/

func main() {
	nums1 := []int{1, 3, 4, 6, 7, 10}
	nums2 := []int{1, 2, 4, 87, 19, 10}
	r := findMedianSortedArrays(nums1, nums2)
	fmt.Println(r)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	aLen := len(nums1)
	bLen := len(nums2)
	allLen := aLen + bLen
	i, j := 0, 0
	var tmpNums []int
	if allLen%2 == 1 {
		for {
			if i < aLen && j < bLen {
				if nums1[i] <= nums2[j] {
					tmpNums = append(tmpNums, nums1[i])
					i++
				} else {
					tmpNums = append(tmpNums, nums2[j])
					j++
				}
			} else if i >= aLen && j < bLen {
				tmpNums = append(tmpNums, nums2[j])
				j++
			} else if j >= bLen && i < aLen {
				tmpNums = append(tmpNums, nums1[i])
				i++
			}

			if len(tmpNums) == allLen/2+1 {
				return float64(tmpNums[allLen/2+1-1])
			}
		}
	}else {
		for {
			if i < aLen && j < bLen {
				if nums1[i] <= nums2[j]{
					tmpNums = append(tmpNums, nums1[i])
					i++
				}else {
					tmpNums = append(tmpNums, nums2[j])
					j++
				}
			} else if i >= aLen && j < bLen{
				tmpNums = append(tmpNums, nums2[j])
				j++
			}else if j >= bLen && i < aLen{
				tmpNums = append(tmpNums, nums1[i])
				i++
			}

			if len(tmpNums) == allLen/2+1{
				return float64((tmpNums[allLen/2+1-1] + tmpNums[allLen/2-1])) / 2.0
			}
		}
	}

	return 0
}

