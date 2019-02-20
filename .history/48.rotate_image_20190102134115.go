/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 48.rotate_image.go
#   Created       : 2019-01-02 11:53:07
#   Describe      :
#
# ====================================================*/
package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix) / 2
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-j][n-1-i]
			matrix[n-1-j][n-1-i] = matrix[i][n-1-j]
			matrix[i][n-1-j] = tmp
		}
	}
	fmt.Println(matrix)
}
