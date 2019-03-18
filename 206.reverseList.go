/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 206.reverseList.go
#   Created       : 2019-03-13 15:05:42
#   Last Modified : 2019-03-13 15:05:42
#   Describe      :
#
# ====================================================*/
package main

func reverseList(head *ListNode) *ListNode {
	var rev *ListNode
	p := head
	for p != nil {
		rev, p, p.Next = p, p.Next, rev
	}
	return rev
}
