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

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	dummy.Next = head

	for pre.Next != nil && pre.Next.Next != nil {
		temp := pre.Next.Next
		pre.Next.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
		pre = temp.Next
	}

	return dummy.Next
}
