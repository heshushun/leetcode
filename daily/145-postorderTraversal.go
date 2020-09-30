package main
/*
145. 二叉树的后序遍历
给定一个二叉树，返回它的 后序 遍历。
示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
*/

func main() {
	
}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Stack []*TreeNode
func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}
func (s *Stack) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
	return n
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, rest := Stack([]*TreeNode{root}), []int{}
	for len(stack) > 0 {
		cur := stack.Pop()
		if cur != nil {
			stack.Push(cur)
			stack.Push(nil)
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
			if cur.Left != nil {
				stack.Push(cur.Left)
			}
		} else {
			rest = append(rest, stack.Pop().Val)
		}
	}
	return rest
}