/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 23.merge_k_sorted_lists.go
#   Created       : 2019-02-12 16:32:47
#   Last Modified : 2019-02-12 16:32:47
#   Describe      :
#
# ====================================================*/
package main

import "sort"

func mergeKLists(lists []*ListNode) *ListNode {
	var res = &ListNode{}

	head := res
	var tempList []int
	for _, v := range lists {
		for v != nil {
			tempList = append(tempList, v.Val)
			v = v.Next
		}
	}

	sort.Ints(tempList)

	for _, v := range tempList {
		head.Next = &ListNode{
			Val: v,
		}
		head = head.Next
	}
	return res.Next
}
