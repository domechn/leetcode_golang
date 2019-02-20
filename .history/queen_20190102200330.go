/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : queen.go
#   Created       : 2019-01-02 19:50:37
#   Describe      :
#
# ====================================================*/
package main

import "fmt"

func main() {
	res := make([]int, 8)
	queen(&res, 0)
}

func queen(res *[]int, row int) {
	if row == 8 {
		fmt.Println(*res)
		return
	}
	for i := 0; i < 8; i++ {
		if isOk(res, row, i) {
			(*res)[row] = i
			queen(res, row+1)
		}
	}
}

func isOk(res *[]int, row, column int) bool {
	leftup := column - 1
	rightup := column + 1
	for i := row - 1; i >= 0; i-- {
		if (*res)[i] == column || ((*res)[i] == leftup && leftup >= 0) || ((*res)[i] == rightup && rightup < 8) {
			return false
		}
		leftup--
		rightup++
	}
	return true
}
