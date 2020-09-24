package main

import "fmt"

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
	tree1 := &TreeNode{1, &TreeNode{3,nil,nil}, &TreeNode{5,nil,nil}}
	r := findMode(tree1)
	print(r)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorder (head *TreeNode){
	if head == nil{
		return
	}
	fmt.Print(head.Val)
	if head.Left != nil {
		preorder(head.Left)
	}
	if head.Right != nil {
		preorder(head.Right)
	}
}
var max int  // 最大计数值
var counter int // 当前计数值
var res []int // 结果
var cur int // 当前值

func findMode(root *TreeNode) []int {
	res, max, cur, counter = []int{}, 1, 0, 0
	dfs(root)
	return res
}

// 深度优先（dfs） 中序遍历（BST值是顺序的）
func dfs(root *TreeNode) {
	if root != nil {
		// 左子树
		dfs(root.Left)

		if root.Val != cur {
			counter = 0
		}
		counter ++
		if max < counter {
			max = counter
			res = []int{root.Val} // 清空
		}else if max == counter {
			res = append(res, root.Val)
		}

		// 当前节点
		cur = root.Val
		// 右子树
		dfs(root.Right)
	}
}