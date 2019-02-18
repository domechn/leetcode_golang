/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 112.path_sum.go
#   Created       : 2019-02-14 10:16:35
#   Last Modified : 2019-02-14 10:16:35
#   Describe      :
#
# ====================================================*/
package leet

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
