/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 102.binary_tree_level_order_traversal.go
#   Created       : 2019-02-26 15:04:46
#   Last Modified : 2019-02-26 15:04:46
#   Describe      :
#
# ====================================================*/
package main

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var temp = []*TreeNode{root}

	for len(temp) != 0 {

		var tt []int

		size := len(temp)
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
		res = append(res, tt)
	}
	return res
}
