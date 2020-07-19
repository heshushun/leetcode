package main

import (
	"fmt"
)

/*
23. 合并K个排序链表
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
func main() {
	l1 := &ListNode{1, &ListNode{2, nil}}
	l2 := &ListNode{2, &ListNode{3, nil}}
	lists := []*ListNode{l1, l2}
	r := mergeKLists(lists)
	for r != nil {
		fmt.Print(r.Val)
		r = r.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

/*将链表拆分为lists[:len(lists)/2]和lists[len(lists)/2:]两部分，
分别使用mergeKLists进行递归直至lists中只有一个链表时返回，
然后返回的结果两两合并，最终合并为一个链表。
*/
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}else if n == 1 {
		return lists[0]
	}
	// 分治法
	return merge(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {

	prehead := &ListNode{}
	result := prehead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		}else {
			prehead.Next =l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}
	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next =l2
	}
	return result.Next
}