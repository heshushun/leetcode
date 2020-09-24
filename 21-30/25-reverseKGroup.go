package main

import (
	"fmt"
)

/*
25. K 个一组翻转链表
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例：
给你这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5
当 k = 3 时，应当返回: 3->2->1->4->5
*/
func main() {
	l := &ListNode{1, &ListNode{2, nil}}
	r := reverseKGroup(l, 2)
	for r != nil {
		fmt.Print(r.Val)
		r = r.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1{
		return head
	}
	// 哨兵
	guard := &ListNode{Next: head}

	// nodek表示第k个节点
	nodek := guard
	for i := 0; i < k; i++ {
		if nodek.Next == nil {
			return guard.Next
		}
		nodek = nodek.Next
	}
	zero := guard

	for {
		// 翻转
		first := zero.Next
		l, r := first, first.Next // 相邻的两个节点（左右节点）
		zero.Next = nodek // 翻转后首连接
		first.Next = nodek.Next // 翻转后尾连接
		for i := 2; i <= k; i++ {
			tmp := r.Next
			r.Next = l // 交换左右两个节点
			l, r = r, tmp
		}

		// 移动到下一个第k节点
		nodek = first
		for i := 0; i < k; i++ {
			if nodek.Next == nil {
				return guard.Next
			}
			nodek = nodek.Next
		}
		zero = first
	}

}