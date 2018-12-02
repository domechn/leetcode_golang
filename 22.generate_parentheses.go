// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func generateParenthesis(n int) []string {
	var res []string
	generateParenthesisDFS(n, n, "", &res)
	return res
}

func generateParenthesisDFS(left, right int, s string, res *[]string) {
	if left > right {
		return
	}

	if left == 0 && right == 0 {
		*res = append(*res, s)
	} else {
		if left > 0 {
			generateParenthesisDFS(left-1, right, s+"(", res)
		}
		if right > 0 {
			generateParenthesisDFS(left, right-1, s+")", res)
		}
	}
}
