/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 107.binary_tree_level_order_traversal_ii.go
#   Created       : 2019-02-26 15:09:56
#   Last Modified : 2019-02-26 15:09:56
#   Describe      :
#
# ====================================================*/
package main

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	temp := []*TreeNode{root}
	for len(temp) != 0 {
		size := len(temp)
		tt := []int{}
		for i := 0; i < size; i++ {
			front := temp[0]
			temp = temp[1:]

			tt = append(tt, front.Val)
			if front.Left != nil {
				temp = append(temp, front.Left)
			}
			if front.Right != nil {
				temp = append(temp, front.Right)
			}

		}
		if len(tt) != 0 {
			res = append([][]int{tt}, res...)
		}
	}
	return res
}
