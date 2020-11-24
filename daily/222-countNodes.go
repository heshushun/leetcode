package main
/*
222. 完全二叉树的节点个数
给出一个完全二叉树，求出该树的节点个数。

说明：
完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

示例:
输入:
    1
   / \
  2   3
 / \  /
4  5 6
输出: 6
*/

func main() {
	tree1 := &TreeNode{1, &TreeNode{3,nil,nil}, &TreeNode{5,nil,nil}}
	r := countNodes(tree1)
	println(r)
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil { // 递归的出口
		return 0
	}
	lHigh, rHigh := 0, 0 // 两侧高度
	lNode, rNode := root, root // 两个指针

	for lNode != nil {  // 计算左侧高度
		lHigh ++
		lNode = lNode.Left
	}

	for rNode != nil {  // 计算右侧高度
		rHigh ++
		rNode = rNode.Right
	}

	if lHigh == rHigh {  // 当前子树是满二叉树，返回节点数
		return 1 << lHigh -1  // 左移N位就是乘以2的n次方
	}

	// 当前子树不是完美二叉树，只是完全二叉树，递归处理左右子树
	return 1 + countNodes(root.Left) + countNodes(root.Right)

}