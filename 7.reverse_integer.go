// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

import (
	"fmt"
)

const Max = 2147483647

func main() {
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	var res int
	var flag bool
	if x < 0 {
		x = -x
		flag = true
	}

	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}

	if flag {
		res = -res
	}
	if res > Max || res < -Max {
		res = 0
	}
	return res
}
