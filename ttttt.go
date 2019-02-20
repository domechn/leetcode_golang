/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : ttttt.go
#   Created       : 2019-01-04 15:20:14
#   Last Modified : 2019-01-04 15:20:14
#   Describe      :
#
# ====================================================*/
package main

import "fmt"

func main() {
	fmt.Println(testt(3, 3, 0, 0))
	fmt.Println(testt(3, 3, 3, 3))
}

func testt(n, w, d1, d2 int) bool {
	if n%3 != 0 || d1 > n/3 || d2 > n/3 {
		return false
	}
	if n == w {
		return d1 == 0 && d2 == 0
	}
	if n == w+1 {
		return (d1 == 1 && d2 == 0) || (d1 == 0 && d2 == 1)
	}
	if n == w+2 {
		return (d1 == 0 && d2 == 2) || (d1 == 2 && d2 == 0) || (d1 == 1 && d2 == 1)
	}
	return false
}
