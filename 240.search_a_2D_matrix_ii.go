/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 240.search_a_2D_matrix_ii.go
#   Created       : 2019-03-04 16:32:56
#   Last Modified : 2019-03-04 16:32:56
#   Describe      :
#
# ====================================================*/
package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i := 0
	j := len(matrix[0]) - 1
	for i < len(matrix) && j >= 0 {

		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] == target {
			return true
		} else {
			i++
		}
	}
	return false
}
