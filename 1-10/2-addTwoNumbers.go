package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}

	l := new(ListNode)
	l.Val = 0
	tmp := l
	flag := 0 //模拟假发中的进位加 1

	for {
		num := l1.Val + l2.Val + tmp.Val
		if num >= 10 {
			tmp.Val = num %10
			flag = 1
		}else {
			tmp.Val = num
		}

		if l1.Next == nil && l2.Next == nil {
			break
		}
		l1 = l1.Next
		l2 = l2.Next
		if l1 == nil {
			l1 = new(ListNode)
			l1.Val = 0
		}
		if l2 == nil {
			l2 = new(ListNode)
			l2.Val = 0
		}

		tmp.Next = new(ListNode)
		if flag == 1 {
			tmp.Next.Val = 1
			flag = 0
		}else {
			tmp.Next.Val = 0
		}
		tmp = tmp.Next
	}
	if flag == 1{
		tmp.Next = new(ListNode)
		tmp.Next.Val = 1
	}
	return l

}

// 返回链表中的值
func (l *ListNode) getInt() (int) {
    if l == nil {
        return 0
    }
    nums := []int{}
    for l != nil {
        nums = append(nums, l.Val)
        l = l.Next
    }
    rst := 0
    for k, v := range nums {
        rst += pow10(k) * v
    }
    return rst
}

// 返回10的几次方 <---- 大数字的情况下在这里首先溢出, 写在这里只为了测试用
func pow10(a int) int {
    if a == 0 {
        return 1
    }
    rst := 1
    for i := 1; i <= a; i++ {
        rst *= 10
    }
    return rst
}

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)

	tmp1 := l1
	tmp2 := l2
	num1 := []int{3, 5, 2}
	num2 := []int{5, 6, 4}

	for i := 0; i < len(num1); i++ {
		tmp1.Val = num1[i]
		if i == len(num1) -1{
			break
		}
		tmp1.Next = new(ListNode)
		tmp1 = tmp1.Next
	}
	for i := 0; i < len(num2); i++ {
		tmp2.Val = num2[i]
		if i == len(num2) -1{
			break
		}
		tmp2.Next = new(ListNode)
		tmp2 = tmp2.Next
	}

	fmt.Println(l2.getInt())
	fmt.Println(l1.getInt())
	l2 = addTwoNumbers(l1, l2)

	fmt.Println(l2.getInt())
}
