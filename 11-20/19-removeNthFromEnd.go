package main

import (
	"fmt"
)

/*
19. 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.

说明：
给定的 n 保证是有效的。
*/

func main() {
	s := ListNode{1, nil}
	r := removeNthFromEnd(&s, 2)
	fmt.Println(r)
}

type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode  {
	m := make(map[int]*ListNode)
	i := 0
	for v:= head; v != nil; v = v.Next {
		m[i] = v
		i ++
	}
	vv, ok := m[i-n]
	if !ok {
		return head
	}
	if i-n-1 < 0{
		return vv.Next
	}
	m[i-n-1].Next = vv.Next
	return head
}