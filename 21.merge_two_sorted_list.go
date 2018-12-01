// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

const MaxInt32 = 1<<31 - 1

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	temp := res
	for l1 != nil || l2 != nil {

		var num1, num2 = MaxInt32, MaxInt32
		if l1 != nil {
			num1 = l1.Val
		}

		if l2 != nil {
			num2 = l2.Val
		}

		if num1 < num2 {
			temp.Next = &ListNode{
				Val: num1,
			}
			l1 = l1.Next
		} else {
			temp.Next = &ListNode{
				Val: num2,
			}
			l2 = l2.Next
		}

		temp = temp.Next

	}

	return res.Next
}
