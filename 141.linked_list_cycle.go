/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 141.linked_list_cycle.go
#   Created       : 2019-02-12 16:12:45
#   Last Modified : 2019-02-12 16:12:45
#   Describe      :
#
# ====================================================*/
package leet

func hasCycle(head *ListNode) bool {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
