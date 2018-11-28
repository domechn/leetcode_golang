// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{
		Val: -1,
	}
	cur := res
	var carry int
	for l1 != nil || l2 != nil {
		var num1, num2 int
		if l1 != nil {
			num1 = l1.Val
		}
		if l2 != nil {
			num2 = l2.Val
		}
		sum := num1 + num2 + carry
		if sum >= 10 {
			carry = 1
			sum %= 10
		} else {
			carry = 0
		}
		cur.Next = &ListNode{
			Val: sum,
		}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry == 1 {
		cur.Next = &ListNode{
			Val: 1,
		}
	}
	return res.Next
}
