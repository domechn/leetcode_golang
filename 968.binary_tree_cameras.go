/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 968.binary_tree_cameras.go
#   Created       : 2019-02-20 10:34:21
#   Last Modified : 2019-02-20 10:34:21
#   Describe      :
#
# ====================================================*/
package leet

func minCameraCover(root *TreeNode) int {
	res := slove(root)
	return min(res[1], res[2])
}

func slove(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0, 9999999}
	}
	L := slove(root.Left)
	R := slove(root.Right)

	minL := min(L[1], L[2])
	minR := min(R[1], R[2])

	d0 := L[1] + R[1]

	d1 := min(L[2]+minR, R[2]+minL)
	d2 := 1 + min(L[0], minL) + min(R[0], minR)
	return []int{d0, d1, d2}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
