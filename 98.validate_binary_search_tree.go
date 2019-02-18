/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 98.validate_binary_search_tree.go
#   Created       : 2019-02-14 10:32:04
#   Last Modified : 2019-02-14 10:32:04
#   Describe      :
#
# ====================================================*/
package leet

const (
	MaxInt = 1<<33 - 1
	MinInt = -1 << 33
)

func isValidBST(root *TreeNode) bool {
	return isValid(root, MinInt, MaxInt)
}

func isValid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
}
