/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 103.binary_tree_zigzag_level_order_traversal.go
#   Created       : 2019-02-26 14:51:13
#   Last Modified : 2019-02-26 14:51:13
#   Describe      :
#
# ====================================================*/
package leet

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res = [][]int{}
	if root == nil {
		return res
	}
	temp := []*TreeNode{root}
	left := true

	for len(temp) != 0 {
		size := len(temp)
		var tt []int
		for i := 0; i < size; i++ {
			tree := front(&temp)

			if tree.Left != nil {
				temp = append(temp, tree.Left)
			}
			if tree.Right != nil {
				temp = append(temp, tree.Right)
			}

			if left {
				tt = append(tt, tree.Val)
			} else {
				tt = append([]int{tree.Val}, tt...)
			}

		}
		left = !left
		if len(tt) != 0 {
			res = append(res, tt)
		}

	}
	return res
}

func front(tree *[]*TreeNode) *TreeNode {
	temp := (*tree)[0]
	*tree = (*tree)[1:]
	return temp
}
