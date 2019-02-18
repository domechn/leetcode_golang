/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 104.maximum_depth_of_binary_tree.go
#   Created       : 2019-02-14 10:40:51
#   Last Modified : 2019-02-14 10:40:51
#   Describe      :
#
# ====================================================*/
package leet

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
