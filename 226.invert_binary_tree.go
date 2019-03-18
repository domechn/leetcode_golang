/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 226.invert_binary_tree.go
#   Created       : 2019-02-14 10:43:56
#   Last Modified : 2019-02-14 10:43:56
#   Describe      :
#
# ====================================================*/
package main

func invertTree(root *TreeNode) *TreeNode {
	temp := root
	invert(temp)
	return root
}

func invert(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	invert(root.Left)
	invert(root.Right)
}
