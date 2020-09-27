package main

/*
235. 二叉搜索树的最近公共祖先
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
示例 1:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。
*/
func main() {
	tree1 := &TreeNode{3, &TreeNode{1,nil,nil}, &TreeNode{5,nil,nil}}
	p := &TreeNode{1, nil, nil}
	q := &TreeNode{2, nil, nil}
	r := lowestCommonAncestor(tree1, p, q)
	print(r)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 因为是二叉搜索树
	if p.Val > q.Val {
		p, q = q, p
	}

	if root == nil || (p.Val <= root.Val && root.Val <= q.Val) {
		return root
	} else if q.Val <= root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}