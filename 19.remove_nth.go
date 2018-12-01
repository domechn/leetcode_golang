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

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	begin := head

	cur := head

	i := 0
	for {
		if i < n {
			cur = cur.Next
			i++
			if cur == nil {
				return head.Next
			}
			continue
		}

		if cur.Next == nil {
			begin.Next = begin.Next.Next
			break
		}

		cur = cur.Next
		begin = begin.Next

	}
	return head
}
