package main

import (
	"fmt"
)

/*
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.
*/
func main() {
	l := &ListNode{1, &ListNode{2, nil}}
	r := swapPairs(l)
	for r != nil {
		fmt.Print(r.Val)
		r = r.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var (
		// 头指针
		h = &ListNode{Val: 0, Next: head}
		// 跟随指针
		p = h
		// 当前节点
		cur = h.Next
	)

	for cur != nil && cur.Next != nil {
		p.Next = cur.Next
		cur.Next = cur.Next.Next
		p.Next.Next = cur

		p = cur
		cur = cur.Next
	}

	return h.Next
}