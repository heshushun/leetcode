package main

import "fmt"

/*
501. 二叉搜索树中的众数
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
假定 BST 有如下定义：
结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树

例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
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