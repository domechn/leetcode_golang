/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 124.maxPathSum.go
#   Created       : 2019-03-13 14:24:20
#   Last Modified : 2019-03-13 14:24:20
#   Describe      :
#
# ====================================================*/
package leet

import "math"

func maxPathSum(root *TreeNode) int {
	var ret = int(math.MinInt32)
	maxRet(&ret, root)
	return ret

}

func maxRet(ret *int, root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0, maxRet(ret, root.Left))
	right := max(0, maxRet(ret, root.Right))

	*ret = max(*ret, root.Val+left+right)
	return max(left, right) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
