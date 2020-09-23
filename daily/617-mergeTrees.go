package main

import (
	"fmt"
)

/*
617. 合并二叉树
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。
示例 1:
输入:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
输出:
合并后的树:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7
注意: 合并必须从两个树的根节点开始。
*/

func main() {
	tree1 := &TreeNode{1, &TreeNode{3,nil,nil}, &TreeNode{5,nil,nil}}
	tree2 := &TreeNode{4, &TreeNode{2,nil,nil}, &TreeNode{6,nil,nil}}
	r := mergeTrees(tree1, tree2)
	for r != nil {
		preorder(r)
	}
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

// 深度优先 递归法
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}
	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
