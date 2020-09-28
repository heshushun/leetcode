package main
/*
117. 填充每个节点的下一个右侧节点指针 II
给定一个二叉树
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有 next 指针都被设置为 NULL。
*/
func main() {
	r := connect(nil)
	println(r)
}

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}

// 树的层次遍历（基于广度优先）
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for i, node := range tmp {
			if i + 1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}