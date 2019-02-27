/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 655.print_binary_tree.go
#   Created       : 2019-02-27 16:33:15
#   Last Modified : 2019-02-27 16:33:15
#   Describe      :
#
# ====================================================*/
package leet

import "strconv"

func printTree(root *TreeNode) [][]string {
	height := getHeight(root)

	res := make([][]string, height)
	p := pow(2, height) - 1
	for i := 0; i < height; i++ {
		res[i] = make([]string, p)
	}

	print(root, &res, height, 0, 0, p-1)
	return res

}

func print(root *TreeNode, res *[][]string, totalHeight, height, i, j int) {
	if root == nil || height == totalHeight {
		return
	}
	(*res)[height][(i+j)/2] = strconv.Itoa(root.Val)
	print(root.Left, res, totalHeight, height+1, i, (j+i)/2-1)
	print(root.Right, res, totalHeight, height+1, (i+j)/2+1, j)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pow(a, b int) int {
	var res = 1
	for i := 0; i < b; i++ {
		res = res * a
	}

	return res
}
